package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/canvas", handleCanvas)
	log.Printf("server start on localhost:%d\n", 18080)
	err := http.ListenAndServe(":18080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func handleCanvas(w http.ResponseWriter, r *http.Request) {
	log.Println("handle canvas...")
	defer r.Body.Close()

	w.Header().Set("Content-Type", "image/svg+xml")

	content, err := read("/tmp/inputfile")
	if err != nil {
		return
	}
	data, err := parse(content)
	if err != nil {
		log.Printf("data check error: %v\n", err)
		return
	}

	err = draw(data, w, 200)
	if err != nil {
		log.Printf("draw data error: %v\n", err)
		return
	}
}
