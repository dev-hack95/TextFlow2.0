package handlers

import (
	"encoding/json"
	"github/dev-hack95/Textflow/controllers"
	"github/dev-hack95/Textflow/structs"
	"github/dev-hack95/Textflow/utilities"
	"github/dev-hack95/Textflow/utilities/logs"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	returnData := utilities.ResponseJson{}
	claims, err := utilities.GetUserSessionDetails(c)
	if err != nil {
		utilities.ErrorResponse(&returnData, "Unauthorized User!")
		return
	}

	if !claims {
		utilities.ErrorResponse(&returnData, "Session Expired!")
		return
	}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		utilities.ErrorResponse(&returnData, "Error occured while reading paylod!")
		return
	}

	inputobj := structs.VideoPayload{}

	err = json.Unmarshal(body, &inputobj)

	if err != nil {
		logs.Error("Error: ", err.Error())
	}

	switch {
	case utilities.IsEmpty(body):
		returnData = controllers.Upload(inputobj)
	default:
		utilities.ErrorResponse(&returnData, "Error occured at uploding data!")
	}

	c.JSON(returnData.Code, returnData)
}

func Produce(c *gin.Context) {
	returnData := utilities.ResponseJson{}
	claims, err := utilities.GetUserSessionDetails(c)
	if err != nil {
		utilities.ErrorResponse(&returnData, "Unauthorized User!")
		return
	}

	if !claims {
		utilities.ErrorResponse(&returnData, "Session Expired!")
		return
	}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		utilities.ErrorResponse(&returnData, "Error occured while reading paylod!")
		return
	}

	inputobj := structs.VideoPayload{}

	err = json.Unmarshal(body, &inputobj)

	if err != nil {
		logs.Error("Error: ", err.Error())
	}

	switch {
	case utilities.IsEmpty(body):
		returnData = controllers.Produce(inputobj)
	default:
		utilities.ErrorResponse(&returnData, "Error occured at uploding data!")
	}

	c.JSON(returnData.Code, returnData)
}
