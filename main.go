package main

import (
	"log"
	"os"

	"github.com/bitcodr/avn-promotion/provider"
)

func main() {

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	if err := os.Setenv("PROMOTION_SERVICE_ROOT_DIR", currentDir); err != nil {
		log.Fatalln(err)
	}

	provider.Init()
}
