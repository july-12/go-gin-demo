package routes

import (
	"starter-with-docker/controller"

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

	r.GET("/posts/:id", controller.PostShow)
	r.GET("/posts", controller.PostIndex)

}

func ProtectRoutes(r *gin.RouterGroup) {

	r.GET("/posts/:id/comments", controller.CommentShow)
	r.POST("/posts/:id/tags", controller.TagsIndexByPost)
	r.POST("/posts/:id/comments", controller.CommentCreate)
	r.POST("/posts", controller.PostCreate)

	r.PUT("/tags/:id", controller.TagUpdate)
	r.POST("/tags", controller.TagCreate)
	r.GET("/tags", controller.TagIndex)

	r.GET("/users/:id/posts", controller.PostIndexOfUser)
	r.GET("/users/:id", controller.UserShow)
	r.GET("/users", controller.UserIndex)
}
