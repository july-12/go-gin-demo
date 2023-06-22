package main

import (
	"net/http"
	database "starter-with-docker/config"
	"starter-with-docker/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDb()

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hahah")

	})
	routes.Init(r)

	r.Run(":3000")
}
