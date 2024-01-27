package base

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.GET("", findAll)
	r.POST("", create)
	r.GET(":base_id", findByID)
	r.PUT(":base_id", update)
	r.DELETE(":base_id", delete)
}