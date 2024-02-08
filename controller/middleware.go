package controller

import (
	"net/http"
	"only-test/utils"

	"github.com/gin-gonic/gin"
)

func (con *Controller) TokenValid(c *gin.Context) {
	token := utils.ExtractToken(c)

	err := utils.VerifyToken(token)
	if err != nil {
		WriteError(c, "Invalid Token", http.StatusUnauthorized, err)
		c.Abort()
		return
	}
	c.Next()
}
