package base

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	ID   string `json:"id"`
	Nome string `json:"nome"`
}

var User Usuario

func create(ctx *gin.Context) {
	err := json.NewDecoder(ctx.Request.Body).Decode(&User)
	if err != nil {
		log.Fatal(error(err))
	}
	ctx.JSON(http.StatusOK,
		gin.H{
			"user": User,
			"method": "POST",
		})
}

func findByID(ctx *gin.Context) {
	userID := ctx.Param("base_id")
	ctx.JSON(http.StatusOK,
		gin.H{
			"base_id": userID,
			"method":     "GET",
		})
}

func findAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK,
		gin.H{
			"method": "GET",
		})
}

func update(ctx *gin.Context) {
	userID := ctx.Param("base_id")
	ctx.JSON(http.StatusOK,
		gin.H{
			"baseID": userID,
			"method": "PUT",
		})
}

func delete(ctx *gin.Context) {
	userID := ctx.Param("base_id")
	ctx.JSON(http.StatusOK,
		gin.H{
			"base_id": userID,
			"method":     "DELETE",
		})
}
