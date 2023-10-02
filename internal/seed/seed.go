package seed

import (
	"context"
	"fmt"
	"log"

	db "github.com/JairoRiver/personal_blog_backend/internal/db/sqlc"
	"github.com/JairoRiver/personal_blog_backend/pkg/util"
	"golang.org/x/term"
)

type Initial struct {
	config util.Config
	store  *db.Queries
}

func New(config util.Config, store *db.Queries) (*Initial, error) {
	initial := Initial{
		config: config,
		store:  store,
	}
	return &initial, nil
}

func (initial *Initial) Run() error {
	//create users admin
	adminUser := createInitialUser(initial, "jairo-admin")
	fmt.Println("user creared", adminUser.Username)

	return nil
}

func createInitialUser(initial *Initial, userName string) db.User {
	ctx := context.Background()

	fmt.Println("Enter a user password please:")
	password, err := term.ReadPassword(0)
	if err != nil {
		log.Fatalln("Can not read password:", err)
	}
	passwordS := string(password)

	var email string
	fmt.Println("Enter an email to user")
	_, err = fmt.Scanln(&email)
	if err != nil {
		log.Fatalln("Can not read email:", err)
	}

	hashedPassword, err := util.HashPassword(passwordS)
	if err != nil {
		log.Fatalln("Error hashed passwoed:", err)
	}

	newUserParams := db.CreateUserParams{
		Username: userName,
		Email:    email,
		Password: hashedPassword,
	}

	user, err := initial.store.CreateUser(ctx, newUserParams)
	if err != nil {
		log.Fatalln("Can not create user:", err)
	}
	log.Println("Created user:", user.Username)

	return user
}
