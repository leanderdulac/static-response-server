package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var contentType = os.Getenv("CONTENT_TYPE")
var responseBody = []byte(os.Getenv("CONTENT_BODY"))

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Static response server listening on port %s.\n", port)
	err := http.ListenAndServe(":"+port, http.HandlerFunc(handler))
	if err != nil {
		panic(err)
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(*http.Request) bool {
		return true
	},
}

func handler(wr http.ResponseWriter, req *http.Request) {
	fmt.Printf("%s | %s %s\n", req.RemoteAddr, req.Method, req.URL)
	if websocket.IsWebSocketUpgrade(req) {
		serveWebSocket(wr, req)
	} else if req.URL.Path == "/.ws" {
		wr.Header().Add("Content-Type", "text/html")
		wr.WriteHeader(200)
		io.WriteString(wr, websocketHTML)
	} else {
		serveHTTP(wr, req)
	}
}

func serveWebSocket(wr http.ResponseWriter, req *http.Request) {
	connection, err := upgrader.Upgrade(wr, req, nil)

	if err != nil {
		fmt.Printf("%s | %s\n", req.RemoteAddr, err)
		return
	}

	defer connection.Close()
	fmt.Printf("%s | upgraded to websocket\n", req.RemoteAddr)

	host, err := os.Hostname()

	if err == nil {
		fmt.Printf("Request served by %s\n", host)
	} else {
		fmt.Printf("Server hostname unknown: %s\n", err.Error())
	}

	err = connection.WriteMessage(websocket.TextMessage, responseBody)

	if err != nil {
		fmt.Printf("%s | %s\n", req.RemoteAddr, err)
	}
}

func serveHTTP(wr http.ResponseWriter, req *http.Request) {
	wr.Header().Add("Content-Type", contentType)
	wr.WriteHeader(200)

	host, err := os.Hostname()
	if err == nil {
		fmt.Printf("Request served by %s\n\n", host)
	} else {
		fmt.Printf("Server hostname unknown: %s\n\n", err.Error())
	}

	wr.Write(responseBody)
}
