package routerHandle

import (
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	param := c.Param("id")
	c.JSON(200, gin.H{
		"ids": param,
	})
}
