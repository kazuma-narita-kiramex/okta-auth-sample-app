package main

import (
	"encoding/json"
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
	log.Print("start hander")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Content-Type", "application/json")

	body := map[string]string{
		"status": "veriry id token ng",
	}

	if err := verifyIdToken(); err == nil {
		body["status"] = "ok"
	}

	bytes, _ := json.Marshal(body)
	w.Write(bytes)
	return
}

func verifyIdToken() error {
	return nil
}
