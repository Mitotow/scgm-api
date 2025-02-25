package main

import (
	"github.com/Mitotow/scgm-api/config"
	"github.com/Mitotow/scgm-api/routers"
	"log"
	"net/http"
	"time"
)

func main() {
	router := routers.CreateRouter()

	server := &http.Server{
		Addr:           ":" + config.GetEnv().AppPort,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
