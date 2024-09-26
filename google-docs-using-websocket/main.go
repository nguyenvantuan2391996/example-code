package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var mapWsConn = make(map[string]*websocket.Conn)

func main() {
	http.HandleFunc("/index", LoadPage)
	http.HandleFunc("/ws", InitWebsocket)
	http.HandleFunc("/save", SaveData)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func LoadPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	path, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	content, err := os.ReadFile(path + "/google-docs-using-websocket/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", content)
}

func InitWebsocket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	channel := r.URL.Query().Get("channel")
	if r.Header.Get("Origin") != "http://"+r.Host {
		http.Error(w, "the origin is invalid", http.StatusInternalServerError)
		return
	}

	if _, ok := mapWsConn[channel]; !ok {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		mapWsConn[channel] = conn
	}
}

func SaveData(w http.ResponseWriter, r *http.Request) {
	channel := r.FormValue("channel")
	data := r.FormValue("data")

	if _, ok := mapWsConn[channel]; !ok {
		http.Error(w, "the channel is not found", http.StatusInternalServerError)
	}

	err := mapWsConn[channel].WriteJSON(map[string]interface{}{
		"data": data,
	})
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("success"))
	if err != nil {
		return
	}
}
