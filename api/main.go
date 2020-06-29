package main

import (
	"github.com/airabinovich/memequotes_front/api/character"
	"github.com/airabinovich/memequotes_front/api/router"
)

func main() {

	character.Initialize()

	engine := router.Route()
	if err := engine.Run(":9001"); err != nil {
		println("Frontend service could not be started")
		panic(err)
	}
}
