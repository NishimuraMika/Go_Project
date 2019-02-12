package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello web server!")
	})

	// ポート = 8000
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("ListenAndServe failed.", err)
	}

}
