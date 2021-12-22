package main

import (
	"log"
	"net/http"
	"time"
	"work/grette_back/database"
	"work/grette_back/routers"
	"work/grette_back/setting"
)

func init() {
	setting.Setup()
	database.Setup()
}

func main() {

	server := &http.Server{
		Addr:    ":8090",
		Handler: routers.InitRouter(),
	}

	log.Print(`[START] Grette START `, time.Now())

	server.ListenAndServe()
}
