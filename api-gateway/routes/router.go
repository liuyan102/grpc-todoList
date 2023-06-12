package routes

import (
	"api-gateway/internal/handler"
	"api-gateway/middlerware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middlerware.Cors(), middlerware.InitMiddleware(service))
	v1 := ginRouter.Group("/api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "success")
		})
		// 用户服务
		v1.POST("/user/register", handler.UserRegister)
		v1.POST("/user/login", handler.UserLogin)

		authed := v1.Group("/")
		authed.Use(middlerware.JWT())
		{
			authed.GET("task", handler.TaskShow)
			authed.POST("task", handler.TaskCreate)
			authed.PUT("task", handler.TaskUpdate)
			authed.DELETE("task", handler.TaskDelete)
		}
	}

	return ginRouter
}
