package utilities

import (
	"net/http"
	"strings"

	"github/dev-hack95/Textflow/helper"

	"github.com/gin-gonic/gin"
)

func GetUserSessionDetails(c *gin.Context) error {
	encodedJwt := c.GetHeader("Authorization")
	encodedJwt = strings.Split(encodedJwt, " ")[1]

	err := helper.VerifyToken(encodedJwt)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Session Expired!"})
		return err
	}

	return nil
}
