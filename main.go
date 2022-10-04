package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nurmanhabib/go-tsa-client/interface/handler"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/digicert-tsa", handler.DigicertTSAHandler)
	http.HandleFunc("/tecxoft-tsa", handler.TecxoftTSAHandler)

	port := os.Getenv("APP_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
