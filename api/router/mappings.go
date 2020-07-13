package router

import (
	"fmt"
	"github.com/airabinovich/memequotes_front/api/character"
	"github.com/airabinovich/memequotes_front/api/config"
	"github.com/airabinovich/memequotes_front/api/phrases"
	"github.com/airabinovich/memequotes_front/api/web"
	"github.com/gin-gonic/gin"
)

func mappings(router *gin.Engine) {

	api := router.Group("api")
	api.GET("characters", character.GetAllCharacters)
	api.GET("character/:character-id/phrases", phrases.GetAllPhrasesForCharacter)

	assets := config.Conf.GetString("web.assets")

	webGroup := router.Group("")
	router.LoadHTMLGlob(fmt.Sprintf("%s/index.html", assets))
	webGroup.Static("/assets", assets)
	webGroup.Static("/static", fmt.Sprintf("%s/static", assets))

	// Index
	webGroup.GET("/", web.IndexHandler)
}
