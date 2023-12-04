package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/surrealdb/surrealdb.go"
    "github.com/adamkali/CharismaSavingThrow/DataService/controllers"
)

func main() {
	// get the environment variables
	godotenv.Load("../.env")
    db, err := surrealdb.New(os.Getenv("CST_DATA_WS"))
    if err != nil {
        panic(err)
    }

    if _, err := db.Signin(map[string]interface{}{
        "user": os.Getenv("CST_DATA_USERNAME"),
        "pass": os.Getenv("CST_DATA_PASSWORD"),
    }); err != nil { panic(err) }

    if _, err := db.Use(
        os.Getenv("CST_DATA_NS"),
        os.Getenv("CST_DATA_DB"),
    ); err != nil { panic(err) }

    var query string
    bs, err := os.ReadFile("./import.surql")
    if err != nil { panic(err) }
    query = string(bs)
    if _, err := db.Query(query, nil); err != nil { panic(err) }

    router := gin.Default()
    dpc := controllers.NewDatePreferenceController(db, router)
    auth := router.Group("/auth")
    {
        datePrefrece := auth.Group("/datePreference")
        {
            datePrefrece.GET("/:dpNumber", dpc.GetDatePreferenceAuth)
            datePrefrece.GET("/", dpc.GetAllDatePreferencesAuth)
        }
    }
    if os.Getenv("CST_USER_DEV") == "true" {
        dev := router.Group("/api/dev")
        {
            datePrefrece := dev.Group("/datePreference")
            {
                datePrefrece.GET("/:dpNumber", dpc.GetDatePreference)
                datePrefrece.GET("/", dpc.GetAllDatePreferences)
            }
        }
    }
    
    router.Run(":" + os.Getenv("CST_DATA_PORT"))
}

