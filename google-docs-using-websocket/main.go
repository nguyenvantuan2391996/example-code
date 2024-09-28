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

var mapWsConn = make(map[string]map[string]*websocket.Conn)

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

	_, err = fmt.Fprintf(w, "%s", content)
	if err != nil {
		return
	}
}

func InitWebsocket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	channel := r.URL.Query().Get("channel")
	uuid := r.URL.Query().Get("uuid")
	if r.Header.Get("Origin") != "http://"+r.Host {
		http.Error(w, "the origin is invalid", http.StatusInternalServerError)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(mapWsConn[channel]) == 0 {
		mapWsConn[channel] = make(map[string]*websocket.Conn)
	}

	mapWsConn[channel][uuid] = conn
}

func SaveData(w http.ResponseWriter, r *http.Request) {
	channel := r.FormValue("channel")
	uuid := r.FormValue("uuid")
	data := r.FormValue("data")

	if _, ok := mapWsConn[channel]; !ok {
		http.Error(w, "the channel is not found", http.StatusInternalServerError)
	}

	for key, ws := range mapWsConn[channel] {
		if key != uuid {
			err := ws.WriteJSON(map[string]interface{}{
				"data": data,
			})
			fmt.Println(websocket.IsCloseError(err))
			fmt.Println(websocket.IsUnexpectedCloseError(err))
			if err != nil {
				return
			}
		}
	}

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("success"))
	if err != nil {
		return
	}
}
