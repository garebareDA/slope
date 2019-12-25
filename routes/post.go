package routes

import(
	"github.com/gin-gonic/gin"
	"fmt"
)

func Post(c *gin.Context) {
	text := c.Query("text")
	token := c.Query("token")

	fmt.Println(text)
	fmt.Println(token)
}