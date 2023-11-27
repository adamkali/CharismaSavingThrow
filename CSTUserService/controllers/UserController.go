package controllers

import (
	"fmt"

	common "github.com/adamkali/CharismaSavingThrow/CSTCommonLib"
	"github.com/adamkali/CharismaSavingThrow/CSTCommonLib/models"
	"github.com/gin-gonic/gin"
	"github.com/surrealdb/surrealdb.go"
)

type UserController struct {
	DB *surrealdb.DB
}

func NewUserController(db *surrealdb.DB) *UserController {
    return &UserController{
        DB: db,
    }
}

// create handles the creation of a new user.
// It binds the JSON request body to a UserRequest struct,
// creates a new user based on the request, and stores it in the database.
// If any errors occur during the process, it returns the appropriate response.
// Finally, it sets the response data to the created user and sends an OK response.
func (c *UserController) create(ctx *gin.Context, dr *common.DetailedResponse) {
    var rep models.UserRequest
    if err := ctx.BindJSON(&rep); err != nil {
        dr.BadRequest(ctx, err)
        return
    }
    user := rep.ToUser()
    if _, err := c.DB.Create(user.GetTableName(), user); err != nil {
        dr.InternalServerError(ctx, err)
        return 
    }

    // get the actual user object rather than the reference
    dr.Data = user
    dr.OK(ctx)
}



// CreateAuth handles the creation of user authentication.
// It validates the HMAC and calls the create function.
func (c *UserController) CreateAuth(ctx *gin.Context) {
    dr := common.NewDetailedResponse(nil)
    if err := common.ValidateHmac(ctx); err != nil {
        dr.Unauthorized(ctx, err)
        return
    }

    c.create(ctx, dr)
}

// Create handles the HTTP POST request to create a new user.
// It calls the create method passing the context and a DetailedResponse object.
/// 
// WARNING: This method is only for development purposes.
func (c *UserController) Create(ctx *gin.Context) {
    dr := common.NewDetailedResponse(nil)
    c.create(ctx, dr)
}

func (c *UserController) get(ctx *gin.Context, dr *common.DetailedResponse) {
    id := ctx.Param("id")
    user := &models.User{}
    userInterface, err := c.DB.Select(id)
    if err != nil {
        dr.InternalServerError(ctx, err)
        return
    }
    if userInterface == nil {
        dr.NotFound(ctx, fmt.Errorf("User with id %s not found", id))
        return
    }
    if err := surrealdb.Unmarshal(userInterface, user); err != nil {
        dr.InternalServerError(ctx, err)
        return
    }
    dr.Data = user
    dr.OK(ctx)
}

func (c *UserController) GetAuth(ctx *gin.Context) {
    dr := common.NewDetailedResponse(nil)
    if err := common.ValidateHmac(ctx); err != nil {
        dr.Unauthorized(ctx, err)
        return
    }

    c.get(ctx, dr)
}

func (c *UserController) Get(ctx *gin.Context) {
    dr := common.NewDetailedResponse(nil)
    c.get(ctx, dr)
}

