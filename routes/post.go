package routes

import (
	"context"
	"strings"
	"log"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"slope/database"
	"time"
	"strconv"
	"errors"
)

type post struct {
	Text  string `json:"text"  binding:"required"`
	Token string `json:"token" binding:"required"`
}

type repryPost struct {
	Text  string `json:"text"  binding:"required"`
	Token string `json:"token" binding:"required"`
	RepryID string `json:"repryID" binding:"required"`
}

//Post クライアントからのPOST受け取り
func Post(c *gin.Context) {
	var posting post
	c.BindJSON(&posting)
	text := posting.Text
	text = strings.Trim(text, " ")

	if text == "" || len(text) > 500 {
		statusError(c, "text error")
	}

	user, err := firebaseUser(posting.Token)
	if err != nil {
		statusError(c, "firebase user error")
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

	photoURL := user.PhotoURL
	if photoURL == "" {
		photoURL = "/static/images/guest.png"
	}

	now := time.Now()
	nowUTC := now.UTC()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	nowJST := nowUTC.In(jst)

	userPost := database.UserPost{}
	userPost.UserName = name
	userPost.Text = text
	userPost.PhotoURL = photoURL
	userPost.CreatedAt = nowJST

	db.Create(&userPost)

	c.JSON(200, gin.H{
		"status": "posted",
	})
}

//RepryPost リプライの受け取り
func RepryPost(c *gin.Context) {
	var repryed repryPost
	c.BindJSON(&repryed)

	text := repryed.Text
	text = strings.Trim(text, " ")

	repryID := repryed.RepryID
	repryID = strings.Trim(repryID, " ")

	id, err := strconv.Atoi(repryID)
	if err != nil {
		statusError(c, "number error")
	}

	if text == "" || len(text) > 500 {
		statusError(c, "text error")
	}

	if id < 0 {
		statusError(c, "id error")
	}

	user, err := firebaseUser(repryed.Token)
	if err != nil {
		statusError(c, "firebase user error")
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

	photoURL := user.PhotoURL
	if photoURL == "" {
		photoURL = "/static/images/guest.png"
	}

	now := time.Now()
	nowUTC := now.UTC()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	nowJST := nowUTC.In(jst)

	repryPost := database.RepryPost{}
	repryPost.UserName = name
	repryPost.Text = text
	repryPost.PhotoURL = photoURL
	repryPost.CreatedAt = nowJST
	repryPost.RepryID = id

	db.Create(&repryPost)
	c.JSON(200, gin.H{
		"status": "posted",
	})
}

func statusError(c *gin.Context, errorMessage string){
	c.JSON(500, gin.H{"status": errorMessage})
	c.Abort()
}

func firebaseUser(tokens string) (*auth.UserRecord, error) {
	ctx := context.Background()
	var err error

	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Println(err)
		err = errors.New("error")
	}

	auth, err := app.Auth(ctx)
	if err != nil {
		log.Println(err)
		err = errors.New("error")
	}

	token, err := auth.VerifyIDToken(ctx, tokens)
	if err != nil {
		log.Println(err)
		err = errors.New("error")
	}

	user, err := auth.GetUser(ctx, token.UID)
	if err != nil {
		log.Println(err)
		err = errors.New("error")
	}

	return user, err
}