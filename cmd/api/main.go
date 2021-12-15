package main

import (
	"github.com/JairoRiver/personal_blog_backend/internal/server"
	"log"
)

func main() {
	server := server.New()

	err := server.Start("127.0.0.1:8080")
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
