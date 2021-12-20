package main

import (
	"work/grette_back/database"
	"work/grette_back/repositories"
	"work/grette_back/setting"
)

func init() {

	setting.Setup()
	database.Setup()
}

func main() {

	com, err := repositories.GetCompanyList()
	_ = com
	_ = err

	// server := &http.Server{
	// 	Addr:    "http://localhost:8090/user/checkNickName",
	// 	Handler: routers.InitRouter(),
	// }

	// log.Print(`[START] Grette START `, time.Now())

	// server.ListenAndServe()
}
