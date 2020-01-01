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
	db, err := database.ConnectDB()
	if err != nil {
		statusError(c, "database error")
	}
	defer db.Close()

	lastPosted := database.UserPost{}
	getNumber, err := strconv.Atoi(get)
	if err != nil {
		statusError(c, "number error")
	}

	firstNum := getNumber
	lastNum := getNumber + 9

	if getNumber < 1 {
		statusError(c, "number error")
	}

	db.Last(&lastPosted)
	lastPostedID := lastPosted.ID

	if lastPostedID < lastNum {
		lastNum = lastPosted.ID
	}

	if lastPostedID < firstNum {
		firstNum = lastPosted.ID
	}

    posts := []database.UserPost{}
	db.Where("created_at BETWEEN ? AND ?", firstNum , lastNum).Find(&posts)
	c.JSON(200, posts)
}