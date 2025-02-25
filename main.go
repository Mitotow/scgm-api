package main

import (
	"github.com/Mitotow/scgm-api/routers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	var err error

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cannot find .env file")
	}

	router := routers.CreateRouter()

	server := &http.Server{
		Addr:           ":" + os.Getenv("APP_PORT"),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
