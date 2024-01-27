package contact

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func create(ctx *gin.Context) {
	// contactID := ctx.Param("contact_id")
	ctx.JSON(http.StatusOK, gin.H{
		// "contact_id": contactID,
		"method": "PUSH",
	})

}
func update(ctx *gin.Context) {
	contactID := ctx.Param("contact_id")
	ctx.JSON(http.StatusOK, gin.H{
		"contact_id": contactID,
		"method": "PUT",
	})

}
func delete(ctx *gin.Context) {
	contactID := ctx.Param("contact_id")
	ctx.JSON(http.StatusOK, gin.H{
		"contact_id": contactID,
		"method": "DELETE",
	})

}
func findByID(ctx *gin.Context) {
	contactID := ctx.Param("contact_id")
	ctx.JSON(http.StatusOK, gin.H{
		"contact_id": contactID,
		"method": "GET",
	})
}
func findAll(ctx *gin.Context) {
	// contactID := ctx.Param("contact_id")
	ctx.JSON(http.StatusOK, gin.H{
		// "contact_id": contactID,
		"method": "GET_ALL",
	})

}
