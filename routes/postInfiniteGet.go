package routes

import(
	"log"
	"strconv"
	"github.com/gin-gonic/gin"
	"slope/database"
)

//PostInfiniteGet infinite load用の関数
func PostInfiniteGet(c *gin.Context) {
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
	lastNum := lastPostedID - (getNumber - 10)

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

//RepryInfiniteGet infinite load repry用
func RepryInfiniteGet(c *gin.Context){
	number := c.Query("number")
	id := c.Query("repry")

	db, err := database.ConnectDB()
	if err != nil {
		statusError(c, "database error")
	}
	defer db.Close()

	getNumber, err := strconv.Atoi(number)
	if err != nil {
		statusError(c, "number error")
	}

	if getNumber < 10 {
		statusError(c, "number error")
	}

	getID, err := strconv.Atoi(id)
	if err != nil {
		statusError(c, "number error")
	}

	lastRepry := database.RepryPost{}
	db.Where("repry_id = ?", getID).Last(&lastRepry)
	lastID := lastRepry.ID

	firstNum := lastID - getNumber
	lastNum := lastID - (getNumber - 10)

	if firstNum < 1 {
		firstNum = 1
	}

	if lastNum < 1 {
		lastNum = 1
	}

	reprys := []database.RepryPost{}
	db.Where("repry_id = ? AND id BETWEEN ? AND ?", getID ,firstNum , lastNum).Find(&reprys)
	c.JSON(200, reprys)
}