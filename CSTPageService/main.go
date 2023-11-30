package main

import (
	controller "github.com/adamkali/CharismaSavingThrow/PageService/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env")
    
    router := gin.Default()
    router.Static("/static", "./static")
    router.LoadHTMLGlob("./routes/*.html")

    router.GET("/", controller.Index)

    user := router.Group("/user")
    {
        user.GET("/check", controller.CheckLoggedIn)
        user.POST("/create", controller.Create)
    }

    
    
}
