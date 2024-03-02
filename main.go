package main

import (
	"hote-booking/pkg/utils"
	"log"
	"os"
)

func main() {
	log.Println("Starting api server")

	configPath := utils.GetConfigPath(os.Getenv("config"))
}
