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
	db.Set("gorm:table_options", "ENGINE = InnoDB CHARSET=utf8mb4",).AutoMigrate(&database.RepryPost{})

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("static", "./static")
	router.GET("/", routes.MainPage)
	router.GET("/posts", routes.PostInfiniteGet)
	router.GET("/posts/post", routes.PostGet)
	router.GET("/posts/reprys", routes.RepryInfiniteGet)

	router.POST("/postText", routes.Post)
	router.POST("/postText/repry", routes.RepryPost)

	router.Run(":8000")
}