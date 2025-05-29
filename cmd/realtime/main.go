package main

import (
	"net/http"
	"snowApp/internal/realtime"

	"log"
)

func main() {
	rtServer := realtime.NewRealtimeServer("localhost:50051")

	http.HandleFunc("/ws", rtServer.HandleConnection)

	log.Println("Realtime server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
