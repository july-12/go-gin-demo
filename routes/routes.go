package routes

import (
	"net/http"
	"starter-with-docker/controller"
	database "starter-with-docker/db"
	"starter-with-docker/models"

	"starter-with-docker/middleware"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	gr := r.Group("api/v1")

	public := gr.Group("/")
	PublicRoutes(public)

	protect := gr.Group("/")
	protect.Use(middleware.AuthorizedRoute)
	ProtectRoutes(protect)

}

func PublicRoutes(r *gin.RouterGroup) {
	r.POST("/signup", controller.Signup)
	r.POST("/login", controller.Login)
	r.POST("/logout", controller.Logout)
}

func ProtectRoutes(r *gin.RouterGroup) {

	r.GET("/facts", func(c *gin.Context) {
		var facts []models.Fact
		database.DB.Find(&facts)
		c.JSON(http.StatusOK, facts)
	})
	r.POST("/fact", func(c *gin.Context) {
		var fact models.Fact
		c.BindJSON(&fact)
		database.DB.Create(&fact)
		var facts []models.Fact
		database.DB.Find(&facts)
		c.JSON(http.StatusOK, facts)
	})

	r.GET("/users", func(c *gin.Context) {
		var users []models.User
		database.DB.Find(&users)
		c.JSON(http.StatusOK, users)
	})

}
