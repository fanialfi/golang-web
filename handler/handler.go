package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	idNumb, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "product page %d", idNumb)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("welcome"))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hallo dunia, saya saat ini sedang belajar golang web"))
}

func MarioHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from nintendo"))
}
