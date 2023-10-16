package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from home!"))
}

func snippetView(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Displaying snippet for %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Request method is not allowed.", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("You are creating a new snippet."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	err := http.ListenAndServe("localhost:4000", mux)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Running on server %d", 4000)
}