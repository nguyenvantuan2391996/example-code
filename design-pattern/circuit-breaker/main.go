package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

type Handler struct {
	CircuitBreaker *CircuitBreaker
}

type ProductData struct {
	ProductName string `json:"product_name"`
	ID          int    `json:"id"`
}

type CircuitBreaker struct {
	State            string
	Wait             time.Duration
	Expiry           int64
	FailureThreshold int64
	Failure          int64
	Mu               sync.Mutex
}

func NewCircuitBreaker(wait time.Duration, failureThreshold int64) *CircuitBreaker {
	return &CircuitBreaker{
		Wait:             wait,
		Mu:               sync.Mutex{},
		State:            Closed,
		FailureThreshold: failureThreshold,
	}
}

const (
	Open     = "open"
	HalfOpen = "half-open"
	Closed   = "closed"
)

func (cb *CircuitBreaker) Allow() bool {
	if cb.Expiry < time.Now().Unix() && cb.State == Open {
		cb.State = HalfOpen
		return true
	}

	if cb.Expiry >= time.Now().Unix() || cb.State == Open {
		return false
	}

	return true
}

func (cb *CircuitBreaker) ResetFailure() {
	cb.Failure = 0
}

func (cb *CircuitBreaker) UpdateState() {
	if cb.Failure > cb.FailureThreshold {
		cb.State = Open
		cb.Expiry = time.Now().Add(cb.Wait).Unix()
		cb.ResetFailure()
	}
}

func (cb *CircuitBreaker) Execute(fc func() ([]byte, error)) (body []byte, err error) {
	if !cb.Allow() {
		fmt.Println("circuit breaker is open")
		return nil, fmt.Errorf("circuit breaker is open")
	}

	switch cb.State {
	case Closed:
		fmt.Println("circuit breaker is closed")
		body, err = fc()
		if err != nil {
			cb.Mu.Lock()
			defer cb.Mu.Unlock()

			cb.Failure++
			cb.UpdateState()
			return nil, err
		}

		return body, err
	case HalfOpen:
		fmt.Println("circuit breaker is half-open")
		cb.Mu.Lock()
		defer cb.Mu.Unlock()

		body, err = fc()
		if err != nil {
			cb.State = Open
			cb.Expiry = time.Now().Add(cb.Wait).Unix()
			return nil, err
		}

		cb.State = Closed
		return body, err
	}

	return
}

func (h *Handler) getExample(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	products, err := h.product()
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		return
	}
}

func (h *Handler) product() ([]*ProductData, error) {
	body, err := h.CircuitBreaker.Execute(func() ([]byte, error) {
		client := &http.Client{}
		req, err := http.NewRequest(http.MethodGet, "http://0.0.0.0:3001/products", nil)

		if err != nil {
			return nil, err
		}

		res, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer func(Body io.ReadCloser) {
			errClose := Body.Close()
			if errClose != nil {
				return
			}
		}(res.Body)

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		return body, nil
	})

	if err != nil {
		return []*ProductData{}, nil
	}

	var response []*ProductData
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func main() {
	h := &Handler{CircuitBreaker: NewCircuitBreaker(5*time.Second, 5)}
	http.HandleFunc("/example", h.getExample)

	// Start the server
	log.Fatal(http.ListenAndServe(":3000", nil))
}
