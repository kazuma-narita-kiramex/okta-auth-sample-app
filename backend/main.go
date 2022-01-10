package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	log.Print("start at :3001")
	if err := http.ListenAndServe(":3001", nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := verifyIdToken(); err != nil {
		fmt.Fprintln(w, "invalid")
	}
	fmt.Fprintln(w, "valid")
}

func verifyIdToken() error {
	return nil
}
