package main

import (
	"fmt"
	"github.com/airabinovich/memequotes_front/api/character"
	"github.com/airabinovich/memequotes_front/api/config"
	"github.com/airabinovich/memequotes_front/api/phrases"
	"github.com/airabinovich/memequotes_front/api/router"
	"log"
	"os"
)

func main() {

	projectDir, err := os.UserHomeDir()
	if err != nil {
		log.Println("ERROR: could not get home directory")
		projectDir = "."
	}
	config.LoadConfiguration(fmt.Sprintf("%s/repos/memequotes_front/api/config/config.dev.conf", projectDir))

	character.Initialize()
	phrases.Initialize()

	engine := router.Route()
	if err := engine.Run(":9001"); err != nil {
		fmt.Println("Frontend service could not be started")
		panic(err)
	}
}
