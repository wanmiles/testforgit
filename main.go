package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Go Server Running Update")
	})

	http.HandleFunc("/update", handleUpdate)

	log.Println("Server started at http://localhost:9999")
	log.Fatal(http.ListenAndServe(":9999", nil))
}
