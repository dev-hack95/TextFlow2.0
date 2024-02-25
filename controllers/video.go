package controllers

import (
	"github/dev-hack95/Textflow/services"
	"github/dev-hack95/Textflow/structs"
	"github/dev-hack95/Textflow/utilities"
)

func Upload(data structs.VideoPayload) (returnData utilities.ResponseJson) {
	//strOuput, err := helper.StringGenerator(data.Video)

	//fmt.Println(strOuput)

	//if err != nil {
	//	utilities.ErrorResponse(&returnData, err.Error())
	//	return
	//}

	//data.Video = strOuput

	response, err := services.UploadVideo(data)

	if err != nil {
		utilities.ErrorResponse(&returnData, err.Error())
		return
	}

	utilities.SuccessResponse(&returnData, response)

	return
}
