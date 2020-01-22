package main

import (
	"io"
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "golang http web test")
	if err != nil {
		log.Fatal("http web test error:", err)
	}
}

func main() {
	http.HandleFunc("/hello", HelloHandler)

	_ = http.ListenAndServe(":10086", nil)
}
