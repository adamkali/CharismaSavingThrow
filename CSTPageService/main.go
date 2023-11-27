package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env")
    
    router := gin.Default()
    router.Static("/static", "./static")
    router.LoadHTMLGlob("./routes/*.html")
    
}
