package main

import (
	"GoRun/configs"
	"GoRun/internal/server"
	"flag"
	"log"
)

var confPath = flag.String("conf-path", "../configs/.env", "Path to config env.")

func main() {
	conf, err := configs.New(*confPath)
	if err != nil {
		log.Fatalln(err)
	}
	server.Run(conf)
}