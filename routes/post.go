package routes

import (
	"context"
	"log"
	"strings"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"slope/database"
)

type post struct {
	Text  string `json:"text"  binding:"required"`
	Token string `json:"token" binding:"required"`
}

//Post クライアントからのPOST受け取り
func Post(c *gin.Context) {
	var posting post
	c.BindJSON(&posting)
	ctx := context.Background()
	text := posting.Text
	text = strings.Trim(text, " ")

	if text == ""{
		statusError(c, "text error")
	}

	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Println(err)
		statusError(c, "firebase down")
	}

	auth, err := app.Auth(ctx)
	if err != nil {
		log.Println(err)
		statusError(c, "firebase down")
	}

	token, err := auth.VerifyIDToken(ctx, posting.Token)
	if err != nil {
		log.Println(err)
		statusError(c, "token error")
	}

	user, err := auth.GetUser(ctx, token.UID)
	if err != nil {
		log.Println(err)
		statusError(c, "token error")
	}

	db, err := database.ConnectDB()
	if err != nil {
		statusError(c, "database error")
	}
	defer db.Close()

	name := user.DisplayName
	if name == "" {
		name = "Guest"
	}

	userPost := database.UserPost{}
	userPost.UserName = name
	userPost.UserUID = user.UID
	userPost.Text = text
	userPost.PhotoURL = user.PhotoURL

	db.Create(&userPost)

	c.JSON(200, gin.H{
		"status": "posted",
	})
}

func statusError(c *gin.Context, errorMessage string){
	c.JSON(500, gin.H{"status": errorMessage})
	c.Abort()
}