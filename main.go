package main

import(
	"github.com/gin-gonic/gin"
	"log"
	"slope/routes"
	"slope/database"
)

func main() {

	db, err := database.ConnectDB()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	db.Set("gorm:table_options", "ENGINE = InnoDB CHARSET=utf8mb4",).AutoMigrate(&database.UserPost{})

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("static", "./static")
	router.GET("/", routes.MainPage)
	router.GET("/posts")

	router.POST("/postText", routes.Post)

	router.Run(":8000")
}