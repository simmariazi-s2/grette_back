package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"work/grette_back/database"
	"work/grette_back/repositories"
	"work/grette_back/routers"
	"work/grette_back/setting"

	_ "gorm.io/gorm"
)

func init() {
	setting.Setup()
	database.Setup()
}

type us struct {
	userNo       int
	userId       string
	userNickname string
	companyNo    int
	userPassword string
}

func main() {

	server := &http.Server{
		Addr:    ":8090",
		Handler: routers.InitRouter(),
	}

	log.Print(`[START] Grette START `, time.Now())

	fmt.Println("db 테스트 진행중 : ", repositories.DbTest())

	server.ListenAndServe()

}
