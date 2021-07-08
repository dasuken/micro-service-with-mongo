package main

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
	"microservices/db"
)

var (
	local bool
)

func init() {
	flag.BoolVar(&local, "local", true, "tun service local")
	flag.Parse()
}

func main() {
	if local {
		// load environment values
		err := godotenv.Load()
		if err != nil {
			log.Panic(err)
		}
	}

	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()
}