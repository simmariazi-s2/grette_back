package main

import (
	"fmt"
	"log"
	"os"
	"work/grette_back/database"
	"work/grette_back/repositories"
	"work/grette_back/setting"

	"github.com/joho/godotenv"
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

	/* server := &http.Server{
		Addr:    ":8090",
		Handler: routers.InitRouter(),
	}

	log.Print(`[START] Grette START `, time.Now())

	server.ListenAndServe() */

	envErr := godotenv.Load(".env")

	if envErr != nil {
		log.Printf(".env file Load Failed")
	}
	dbName := os.Getenv("DB_NAME")
	rootID := os.Getenv("DB_ID")
	rootPW := os.Getenv("DB_PW")
	dbAddr := os.Getenv("DB_ADDR")

	fmt.Println("dbName ::: ", dbName)
	fmt.Println("rootID ::: ", rootID)
	fmt.Println("rootPW ::: ", rootPW)
	fmt.Println("dbAddr ::: ", dbAddr)

	fmt.Println("db 테스트 진행중 : ", repositories.DbTest())
	var a us

	//database.Db.Raw("SELECT * FROM user").Scan(&a)
	database.Db.Select(&a)

	fmt.Println(a.companyNo)

}
