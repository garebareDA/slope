package routes

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

//MainPage メインページの表示
func MainPage(c *gin.Context) {
	c.HTML(http.StatusFound, "index.html", gin.H{})
}