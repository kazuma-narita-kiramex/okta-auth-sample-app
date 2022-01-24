package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type Handler struct {
	JWKS_URL string
	ClientID string
	Issure   string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	handler := Handler{
		JWKS_URL: os.Getenv("JWKS_URL"),
		ClientID: os.Getenv("CLIENT_ID"),
		Issure:   os.Getenv("ISSURE"),
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

	refreshInterval := time.Hour
	options := keyfunc.Options{
		RefreshInterval: refreshInterval,
		RefreshErrorHandler: func(err error) {
			log.Printf("There was an error with the jwt.KeyFunc\nError: %s", err.Error())
		},
	}

	jwks, err := keyfunc.Get(h.JWKS_URL, options)
	if err != nil {
		log.Println("Failed to create JWKS from resource at the given URL.")
		return nil
	}

	// Parse the JWT.
	token, err := jwt.Parse(bearerToken, jwks.Keyfunc)
	if err != nil {
		log.Println("Failed to parse the JWT.")
		return err
	}

	// Check if the token is valid.
	if !token.Valid {
		return errors.New("The token is not valid.")
	}

	// Check Claims
	claims := token.Claims.(jwt.MapClaims)
	if aud, ok := claims["aud"]; ok == false || aud != h.ClientID {
		return errors.New("aud is not valid.")
	}
	if iss, ok := claims["iss"]; ok == false || iss != h.Issure {
		return errors.New("iss is not valid.")
	}

	return nil
}
