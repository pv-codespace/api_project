package routes

import (
	"fmt"
	"net/http"
	"rest_api/models"
	"rest_api/utils"

	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error while binding data to json"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("error while creating user %v", err)})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error while binding data to json"})
		return
	}

	err = user.ValidateUser()

	fmt.Println("error", err)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	fmt.Printf("token err %v", err)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "login successfully!", token: token})
}

// func getUsers(context *gin.Context){

// }
