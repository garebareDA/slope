package routes
import(
	"strconv"
	"github.com/gin-gonic/gin"
	"slope/database"
)

//PostGet 一つのpost
func PostGet(c *gin.Context) {
	get := c.Query("number")
	db, err := database.ConnectDB()
	if err != nil {
		statusError(c, "database error")
	}
	defer db.Close()

	getNumber, err := strconv.Atoi(get)
	if err != nil {
		statusError(c, "number error")
	}

	if getNumber < 0 {
		statusError(c, "number error")
	}

	post := database.UserPost{}
	notfound := db.Where("id = ?", getNumber).First(&post).RecordNotFound()
	if notfound {
		statusError(c, "number error")
	}

	c.JSON(200, post)
}