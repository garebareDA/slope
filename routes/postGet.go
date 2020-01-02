package routes

import(
	"log"
	"strconv"
	"github.com/gin-gonic/gin"
	"slope/database"
)

func PostGet(c *gin.Context) {
	get := c.Query("number")

	log.Println(get)
	db, err := database.ConnectDB()
	if err != nil {
		statusError(c, "database error")
	}
	defer db.Close()

	getNumber, err := strconv.Atoi(get)
	if err != nil {
		statusError(c, "number error")
	}

	if getNumber < 10 {
		statusError(c, "number error")
	}

	lastPosted := database.UserPost{}
	db.Last(&lastPosted)

	lastPostedID := lastPosted.ID
	firstNum := lastPostedID - getNumber
	lastNum := lastPostedID - (getNumber - 9)

	log.Println(firstNum)
	log.Println(lastNum)

	if firstNum < 1 {
		firstNum = 1
	}

	if lastNum < 1 {
		lastNum = 1
	}

	log.Println(firstNum)
	log.Println(lastNum)

    posts := []database.UserPost{}
	db.Where("id BETWEEN ? AND ?", firstNum , lastNum).Find(&posts)
	c.JSON(200, posts)
}