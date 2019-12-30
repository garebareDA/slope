package routes

import(
	"log"
	"strconv"
	"github.com/gin-gonic/gin"
	"slope/database"
)

func postGet(c *gin.Context) {
	get := c.Query("number")
	log.Println(get)

	getNumber, err := strconv.Atoi(get)
	if err != nil {

	}

	if getNumber < 1 {

	}

	db, err := database.ConnectDB()
	if err != nil {
		statusError(c, "database error")
	}
	defer db.Close()

    posts := []database.UserPost{}
	db.Where("created_at BETWEEN ? AND ?", getNumber , getNumber + 9).Find(&posts)
}