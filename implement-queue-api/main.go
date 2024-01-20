package main

import (
	"log"
	"net/http"
)

type Handler struct {
	QueueStruct *Queue
}

func NewHandler(queue *Queue) *Handler {
	return &Handler{
		QueueStruct: queue,
	}
}

func (h *Handler) handleRequestAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// push queue
	h.QueueStruct.Enqueue(r.URL.Query().Get("msg"))
}

func main() {
	queue := NewQueue(NumberOfCapacity)
	go func() {
		queue.Dequeue()
	}()

	handler := NewHandler(queue)

	http.HandleFunc("/push-queue", handler.handleRequestAPI)

	// Start the server
	log.Fatal(http.ListenAndServe(":3000", nil))
}
