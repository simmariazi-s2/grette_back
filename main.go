package main

import (
	database "work/grette_back/database"

	model "work/grette_back/model"
)

func main() {
	clientDB := database.ConnectionDB()
	user := new(model.User)
	_ = clientDB
	_ = user

}
