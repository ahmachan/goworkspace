package router

import (
	"dmhi/handler/users"
	"dmhi/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	aug := g.Group("/users")
	{
		aug.GET("/info", user.GetUserInfo)
		aug.GET("/health", user.HealthCheck)
		aug.GET("/secure",user.DisMac)
		aug.POST("/ak",user.GetAkWithToken)
		aug.GET("/akbucket",user.GetUpHostByAkBucket)
	}

	return g
}
