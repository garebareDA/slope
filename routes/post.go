package routes

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
)

type post struct {
	Text  string `json:"text"  binding:"required"`
	Token string `json:"token" binding:"required"`
}

func Post(c *gin.Context) {
	var posting post
	c.BindJSON(&posting)
	ctx := context.Background()

	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Println(err)
	}

	auth, err := app.Auth(ctx)
	if err != nil {
		log.Println(err)
	}

	token, err := auth.VerifyIDToken(ctx, posting.Token)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}

	log.Println(token)
}
