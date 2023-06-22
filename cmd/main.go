package main

import (
	"log"
	database "starter-with-docker/db"
	"starter-with-docker/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.ConnectDb()

	r := gin.Default()

	routes.Init(r)

	r.Run(":3000")
}
