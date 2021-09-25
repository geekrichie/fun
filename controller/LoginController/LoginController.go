package LoginController

import (
	"net/http"
	"web/model"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	var user = model.NewUser(name, password)
	err := user.Create()
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "创建用户失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "创建用户成功",
		})
	}
}
