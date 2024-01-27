package user

import (
	"github.com/alisson-paulino7/apigin/interfaces/http/user/base"
	"github.com/alisson-paulino7/apigin/interfaces/http/user/contact"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	base.Router(r.Group("base"))
	contact.Router(r.Group("contact"))
}
