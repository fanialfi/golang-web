package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)

	log.Println("starting web on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err.Error())
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hallo dunia, saya saat ini sedang belajar golang web"))
}
