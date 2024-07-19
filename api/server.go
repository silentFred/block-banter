package api

import (
	"log"
	"net/http"
)

func StartWebServer() {
	http.HandleFunc("/api/events", ServeTransferEvents)
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./public"))))

	log.Println("Server is running on port 9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}
