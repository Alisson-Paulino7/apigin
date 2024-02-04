package base

import (
	// "encoding/json"
	// "log"
	"net/http"

	"github.com/alisson-paulino7/apigin/application/user/base"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// type Usuario struct {
// 	ID   string `json:"id"`
// 	Nome string `json:"nome"`
// }

// var User Usuario

func create(ctx *gin.Context) {
	// err := ctx.BindJSON(&User)
	// if err != nil {
	// 	log.Fatal(error(err))
	// }
	// ctx.JSON(http.StatusOK,
	// 	gin.H{
	// 		"user": User,
	// 		"method": "POST",
	// 	})

	var request base.CreateRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	response, err := base.Create(ctx, &request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusCreated, response)
}

func findByID(ctx *gin.Context) {

	userID := ctx.Param("user_id")

	NewUserID, err := uuid.Parse(userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	response ,err := base.FindByID(ctx, &NewUserID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, response)

	// userID := ctx.Param("user_id")
	// ctx.JSON(http.StatusOK,
	// 	gin.H{
	// 		"user_id": userID,
	// 		"method":  "GET",
	// 	})
}

func findAll(ctx *gin.Context) {

	response, err := base.FindAll(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, response)
}

func update(ctx *gin.Context) {

	var request base.UpdateRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	UserID := ctx.Param("user_id")
	NewUserID, err := uuid.Parse(UserID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	request.ID = &NewUserID

	if err := base.Update(ctx, &request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusNoContent, nil)
	// userID := ctx.Param("user_id")
	// ctx.JSON(http.StatusOK,
	// 	gin.H{
	// 		"UserID": userID,
	// 		"method": "PUT",
	// 	})
}

func delete(ctx *gin.Context) {
	userID := ctx.Param("user_id")

	NewUserID, err := uuid.Parse(userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if err := base.Delete(ctx, &NewUserID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusNoContent, nil)

	// ctx.JSON(http.StatusOK,
	// 	gin.H{
	// 		"User_id": userID,
	// 		"method":  "DELETE",
	// 	})
}
