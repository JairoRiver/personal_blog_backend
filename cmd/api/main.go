package main

import (
	"database/sql"
	"github.com/JairoRiver/personal_blog_backend/internal/server"
	"github.com/JairoRiver/personal_blog_backend/internal/util"
	db "github.com/JairoRiver/personal_blog_backend/pkg/db/sqlc"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".", "app")
	if err != nil {
		log.Fatal("can not load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can not connect to db:", err)
	}

	store := db.New(conn)

	server := server.New(config, store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
