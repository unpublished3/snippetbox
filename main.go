package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err  := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 0 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Hello %d from snippedView", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from snippetCreate"))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a snippet"))
}


func main()  {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("starting server on: 4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}