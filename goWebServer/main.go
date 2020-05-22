package main

import "net/http"

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Internet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	httpServer := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}
	httpServer.ListenAndServe()
}
