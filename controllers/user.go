package controllers

import (
	"github/dev-hack95/Textflow/helper"
	"github/dev-hack95/Textflow/models"
	"github/dev-hack95/Textflow/structs"
	"github/dev-hack95/Textflow/utilities"
	"time"
)

func SignUp(data structs.UserSignUp) (returnData utilities.ResponseJson) {
	User, errEmail := models.GetUserByEmail(data.Email)
	if errEmail != nil {
		utilities.ErrorResponse(&returnData, "Error Ocuured while fetching data from database!")
		return
	}
	if User != nil && User.Email == data.Email {
		utilities.ErrorResponse(&returnData, "User is already present in database!")
		return
	}
	token, errToken := helper.CreateToken(data.FirstName, data.LastName, data.Email)
	if errToken != nil {
		utilities.ErrorResponse(&returnData, "Error occured at creating token!")
		return
	}

	password, errPassword := helper.HashPassword(data.Password)
	if errPassword != nil {
		utilities.ErrorResponse(&returnData, "Unable to hash password")
		return
	}

	if User == nil {
		id, err := models.AddUserDetails(
			&models.Users{
				FirstName: data.FirstName,
				LastName:  data.LastName,
				Email:     data.Email,
				Password:  password,
				UserToken: token,
				CreatedAt: time.Now(),
				UpdateAt:  time.Now(),
			})
		if err != nil {
			utilities.ErrorResponse(&returnData, "Error Occured while adding user!")
			return
		}
		utilities.SuccessResponse(&returnData, id)
	}

	return
}

func SignIn(data structs.UserSignIn) (returnData utilities.ResponseJson) {
	User, err := models.GetUserByEmail(data.Email)
	if err != nil {
		utilities.ErrorResponse(&returnData, "Error occured at reading data from database!")
		return
	}
	if User == nil {
		utilities.ErrorResponse(&returnData, "User is not present in database please create account")
		return
	}
	check := helper.VerifyPassword(User.Password, data.Password)
	if User.Email == data.Email && check {
		utilities.SuccessResponse(&returnData, User.UserToken)
		return
	} else {
		utilities.ErrorResponse(&returnData, "Email and Passwod does not match")
	}
	return
}
