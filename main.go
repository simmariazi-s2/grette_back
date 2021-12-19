package main

import (
	"log"
	"net/http"
	"time"
	"work/grette_back/database"
	"work/grette_back/routers"
)

func main() {

	db, err := database.Setup()

	/*
		var user User
		aa := db.Select("email").Where(User{Email: "baekhk1006@gmail.com"}).First(&user)
		_ = aa
	*/
	_ = db
	_ = err
	server := &http.Server{
		Addr:    ":8090",
		Handler: routers.InitRouter(),
	}

	log.Print(`[START] Grette START `, time.Now())

	server.ListenAndServe()
}
