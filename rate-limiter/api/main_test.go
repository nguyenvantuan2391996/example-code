package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"testing"
)

func TestHandlerAPI_testHandler(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			callAPI()
		}()
	}

	wg.Wait()
	fmt.Println("finish")
}

func callAPI() {
	url := "http://localhost:3000/test"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	fmt.Println(res.StatusCode)
}
