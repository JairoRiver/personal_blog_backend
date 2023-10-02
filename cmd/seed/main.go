package main

import (
	"context"
	"log"

	db "github.com/JairoRiver/personal_blog_backend/internal/db/sqlc"
	"github.com/JairoRiver/personal_blog_backend/internal/seed"
	"github.com/JairoRiver/personal_blog_backend/pkg/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config, err := util.LoadConfig(".", "app")
	if err != nil {
		log.Fatal("can not load config:", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db")
	}

	store := db.New(connPool)

	initial, err := seed.New(config, store)
	if err != nil {
		log.Fatal(err)
	}

	initial.Run()
}
