package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/surrealdb/surrealdb.go"
    "github.com/adamkali/CharismaSavingThrow/UserService/controllers"
)

func main() {
	// get the environment variables
	godotenv.Load("../.env")
    println(os.Getenv("CST_USER_WS"))
    db, err := surrealdb.New(os.Getenv("CST_USER_WS"))
    if err != nil {
        panic(err)
    }

    if _, err := db.Signin(map[string]interface{}{
        "user": os.Getenv("CST_USER_USERNAME"),
        "pass": os.Getenv("CST_USER_PASSWORD"),
    }); err != nil { panic(err) }

    if _, err := db.Use(
        os.Getenv("CST_USER_NS"),
        os.Getenv("CST_USER_DB"),
    ); err != nil { panic(err) }
    
    uc := controllers.NewUserController(db)
    
    router := gin.Default()
    auth := router.Group("/api/auth")
    {
        user := auth.Group("/user")
        {
            user.POST("/", uc.CreateAuth)
            user.GET("/:id", uc.GetAuth)
        }
    }
    if os.Getenv("CST_USER_DEV") == "true" {
        dev := router.Group("/api/dev")
        {
            user := dev.Group("/user")
            {
                user.POST("/", uc.Create)
                user.GET("/:id", uc.Get)
            }
        }
    }
    router.Run(":" + os.Getenv("CST_USER_PORT"))
}
