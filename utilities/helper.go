package utilities

import (
	"github/dev-hack95/Textflow/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetUserSessionDetails(c *gin.Context) (bool, error) {
	encodedJwt := c.GetHeader("Authorization")
	encodedJwt = strings.Split(encodedJwt, " ")[1]

	claims, err := helper.VerifyToken(encodedJwt)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Session Expired!"})
		return false, err
	}

	return claims, nil
}
