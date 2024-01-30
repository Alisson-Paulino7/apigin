package base

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.GET("", findAll)
	r.POST("", create)
	r.GET(":user_id", findByID)
	r.PUT(":user_id", update)
	r.DELETE(":user_id", delete)
}