package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	// Ana endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := Response{
			Message: "Hello, World!",
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// Özel mesaj endpoint'i
	http.HandleFunc("/hello/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Path[len("/hello/"):]
		if name == "" {
			name = "World"
		}

		response := Response{
			Message: "Merhaba, " + name + "!",
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	log.Println("Server 8080 portunda başlatılıyor...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}