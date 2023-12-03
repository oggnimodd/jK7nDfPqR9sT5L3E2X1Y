package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oggnimodd/jK7nDfPqR9sT5L3E2X1Y/controllers"
	"github.com/oggnimodd/jK7nDfPqR9sT5L3E2X1Y/middleware"
)

type PostRouteController struct {
	postController controllers.PostController
}

func NewRoutePostController(postController controllers.PostController) PostRouteController {
	return PostRouteController{postController}
}

func (pc *PostRouteController) PostRoute(rg *gin.RouterGroup) {

	router := rg.Group("posts")
	router.Use(middleware.DeserializeUser())
	router.POST("/", pc.postController.CreatePost)
	router.GET("/", pc.postController.FindPosts)
	router.PUT("/:postId", pc.postController.UpdatePost)
	router.GET("/:postId", pc.postController.FindPostById)
	router.DELETE("/:postId", pc.postController.DeletePost)
}
