package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	verifier "github.com/okta/okta-jwt-verifier-golang"
)

type Handler struct {
	OauthIssuer string
	ClientID    string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	handler := Handler{
		OauthIssuer: os.Getenv("OAUTH_ISSUER"),
		ClientID:    os.Getenv("CLIENT_ID"),
	}

	http.HandleFunc("/", handler.index)

	log.Print("start at :3001")
	if err := http.ListenAndServe(":3001", nil); err != nil {
		log.Fatal(err)
	}
}

func (h Handler) index(w http.ResponseWriter, r *http.Request) {
	log.Print("start hander")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")

	if r.Method == http.MethodOptions {
		return
	}

	body := map[string]string{
		"status": "ng",
	}

	authHeader := r.Header.Get("Authorization")
	if err := h.verifyToken(authHeader); err != nil {
		log.Print(err)
	} else {
		body["status"] = "ok"
	}

	w.Header().Set("Content-Type", "application/json")
	bytes, _ := json.Marshal(body)
	w.Write(bytes)
	return
}

func (h Handler) verifyToken(authHeader string) error {
	if authHeader == "" {
		return errors.New("missing id authHeader")
	}
	tokenParts := strings.Split(authHeader, "Bearer ")
	bearerToken := tokenParts[1]

	tv := map[string]string{}
	tv["aud"] = "api://default"
	tv["cid"] = h.ClientID
	jv := verifier.JwtVerifier{
		Issuer:           h.OauthIssuer,
		ClaimsToValidate: tv,
	}

	_, err := jv.New().VerifyAccessToken(bearerToken)
	if err != nil {
		return err
	}

	return nil
}
