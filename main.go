package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/mario", marioHandler)

	log.Println("starting web on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err.Error())
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("welcome"))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hallo dunia, saya saat ini sedang belajar golang web"))
}

func marioHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from nintendo"))
}
