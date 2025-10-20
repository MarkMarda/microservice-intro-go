package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MarkMarda/broker-service/internal/api"
)

const webPort = "80"

func main() {
	app := api.Config{}

	log.Printf("Starting broker service on port %s", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.Routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
