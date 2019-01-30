package main

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/dreadatour/drone-station/api"
)

func main() {
	dotenv := flag.Bool("dotenv", false, "load config from '.env' file")
	flag.Parse()

	if dotenv != nil && *dotenv {
		// load config from '.env' file
		if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
			log.Fatalln(err)
		}
	}

	if err := api.Run(); err != nil {
		log.Fatalln(err)
	}
}
