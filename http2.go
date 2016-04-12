package main

import "net/http"

func main() {
	port := ":8899"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/plain")
		w.Write([]byte(r.Proto))
	})

	http.ListenAndServeTLS(port, "crt/server.crt", "crt/server.key", nil)
}
