package main

import (
	"github.com/JairoRiver/personal_blog_backend/internal/server"
	"github.com/JairoRiver/personal_blog_backend/internal/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".", "app")
	if err != nil {
		log.Fatal("can not load config:", err)
	}

	server := server.New()

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
