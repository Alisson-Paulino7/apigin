package main

import (
	"fmt"
	"log"

	"github.com/alisson-paulino7/apigin/config"
	"github.com/alisson-paulino7/apigin/database"
	"github.com/alisson-paulino7/apigin/interfaces/http/user"
	"github.com/gin-gonic/gin"
)


func main() {

	if err := config.Load(); err != nil {
		log.Fatal(err)
	}

	if err := database.Init(); err != nil {
		log.Fatal(err)
	}

	engine := gin.Default()

	user.Router(engine.Group("users"))

	engine.Run(fmt.Sprintf(":%d", config.Get().HTTPPort)) // listen and serve on 0.0.0.0:8080 as default
}
