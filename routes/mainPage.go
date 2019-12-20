package routes

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

func MainPage(c *gin.Context) {
	c.HTML(http.StatusFound, "index.html", gin.H{})
}