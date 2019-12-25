package main

import(
	"github.com/gin-gonic/gin"
	"slope/routes"
)

//firebaseのAuth認証を使う
//サーバー側ではクライアントから送られてきたデータをSQLに
//アカウントは保持しない
//firebase adminの追加
//uidはTokenで取得

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("static", "./static")
	router.GET("/", routes.MainPage)
	router.POST("/postText", routes.Post)

	router.Run(":8000")
}