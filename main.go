package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	// Register the routes and handlers
	mux.Handle("/", &homeHandler{})

	// Run the server
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Andrade Chris Decierdo"))
	if err != nil {
		return
	}
}
