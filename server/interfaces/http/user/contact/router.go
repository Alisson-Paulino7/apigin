package contact

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.GET("", findAll)
	r.POST("", create)
	r.GET(":contact_id", findByID)
	r.PUT(":contact_id", update)
	r.DELETE(":contact_id", delete)
}
