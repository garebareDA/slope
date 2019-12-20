package main

import(
	"github.com/gin-gonic/gin"
	"slope/routes"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", routes.MainPage)

	router.Run(":8000");
}