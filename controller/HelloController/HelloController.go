package HelloController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "gorm.io/gorm"
)

// @Description hello
// @Accept json
// @Produce json
// @Success 200 {object} model.User
// @Router /hello [get]
func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "hello world"})
}
