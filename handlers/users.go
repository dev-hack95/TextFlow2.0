package handlers

import (
	"encoding/json"
	"github/dev-hack95/Textflow/controllers"
	"github/dev-hack95/Textflow/structs"
	"github/dev-hack95/Textflow/utilities"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	returnData := utilities.ResponseJson{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		utilities.ErrorResponse(&returnData, "Failure: incorrect posts data")
		c.JSON(returnData.Code, returnData)
		return
	}

	inputobj := structs.UserSignUp{}
	err = json.Unmarshal(body, &inputobj)
	if err != nil {
		utilities.ErrorResponse(&returnData, "Failed to unmarshal JSON")
		c.JSON(returnData.Code, returnData)
		return
	}

	switch {
	case !utilities.IsEmpty(inputobj):
		returnData = controllers.SignUp(inputobj)
	default:
		utilities.ErrorResponse(&returnData, "Something Went Wrong")
	}

	c.JSON(returnData.Code, returnData)
}

func SignIn(c *gin.Context) {
	returnData := utilities.ResponseJson{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		utilities.ErrorResponse(&returnData, "Failure: Unable to read body")
		c.JSON(returnData.Code, returnData)
		return
	}

	inputobj := structs.UserSignIn{}
	_ = json.Unmarshal(body, &inputobj)

	switch {
	case !utilities.IsEmpty(inputobj):
		returnData = controllers.SignIn(inputobj)
	default:
		utilities.ErrorResponse(&returnData, "Unbale to login")
	}

	c.JSON(returnData.Code, returnData)
}
