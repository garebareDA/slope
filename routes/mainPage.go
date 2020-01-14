package routes

import(
	"github.com/gin-gonic/gin"
)

//MainPage メインページの表示
func MainPage(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}