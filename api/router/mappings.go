package router

import (
	"github.com/airabinovich/memequotes_front/api/character"
	"github.com/airabinovich/memequotes_front/api/web"
	"github.com/gin-gonic/gin"
)

func mappings(router *gin.Engine) {

	webGroup := router.Group("")
	router.LoadHTMLGlob("/home/ariel/repos/memequotes_front/api/assets/index.html")
	webGroup.Static("/assets", "/home/ariel/repos/memequotes_front/api/assets")
	webGroup.Static("/static", "/home/ariel/repos/memequotes_front/api/assets/static")

	// Index
	webGroup.GET("/", web.IndexHandler)

	api := router.Group("api/")
	api.GET("characters", character.GetAllCharacters)
}
