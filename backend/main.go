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
		fmt.Fprintln(w, "invalid token")
	}
	fmt.Fprintln(w, "valid token")
}

func verifyIdToken() error {
	return nil
}
