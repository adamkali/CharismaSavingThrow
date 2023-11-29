package controllers

import (
	"fmt"
	"strconv"

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

// get handles the retrieval of a user.
// It gets the user id from the request parameters,
// gets the user from the database, and unmarshals it into a User object.
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

// GetAuth handles the retrieval of a user.
// It validates the HMAC and calls the get function.
func (c *UserController) GetAuth(ctx *gin.Context) {
    dr := common.NewDetailedResponse(nil)
    if err := common.ValidateHmac(ctx); err != nil {
        dr.Unauthorized(ctx, err)
        return
    }

    c.get(ctx, dr)
}

// Get handles the HTTP GET request to retrieve a user.
// It calls the get method passing the context and a DetailedResponse object.
//
// WARNING: This method is only for development purposes.   
func (c *UserController) Get(ctx *gin.Context) {
    dr := common.NewDetailedResponse(nil)
    c.get(ctx, dr)
}

// Login handles the login of a user.
// by takin the username and password from the request body,
// creating a hash of the username and password, and then
// comparing it to the hash stored in the database.
// If the hashes match, it returns the user object.
func (c *UserController) login(ctx *gin.Context, dr *common.DetailedResponse) {
    var rep models.UserLoginRequest
    user := &models.User{}
    if err := ctx.BindJSON(&rep); err != nil {
        dr.BadRequest(ctx, err)
        return
    }
    userInterface, err := c.DB.Select(user.GetTableName() + ":" + rep.Username);
    if err != nil {
        dr.NotFound(ctx, err)
        return 
    }
    fmt.Printf("%v", userInterface)
    if err := surrealdb.Unmarshal(userInterface, user); err != nil {
        dr.InternalServerError(ctx, err)
        return
    }
    if !user.ValidateHash(rep.Password) {
        dr.Unauthorized(ctx, fmt.Errorf("Invalid username or password"))
        return
    }
    dr.Data = user
    dr.OK(ctx)
}

// LoginAuth handles the login of a user.
// It validates the HMAC and calls the login function.
func (c *UserController) LoginAuth(ctx *gin.Context) {
    dr := common.NewDetailedResponse(nil)
    if err := common.ValidateHmac(ctx); err != nil {
        dr.Unauthorized(ctx, err)
        return
    }

    c.login(ctx, dr)
}

// Login handles the HTTP POST request to login a user.
// It calls the login method passing the context and a DetailedResponse object.
//
// WARNING: This method is only for development purposes.
func (c *UserController) Login(ctx *gin.Context) {
    dr := common.NewDetailedResponse(nil)
    c.login(ctx, dr)
}

// updateDatePrefrence handles the updating of the users 
// date preference, which is either: 
// - 0: Seeking a relationship
// - 1: Seeking a friendship
// - 2: Seeking both a relationship and a friendship
// - 3: Seeking board game partners
// - 4: Seeking video game partners
// - 5: Seeking anything
func (c *UserController) updateDatePrefrence(ctx *gin.Context, dr *common.DetailedResponse) {
    // get the user id from the request parameters
    id := ctx.Param("id")
    // get the date preference from the request parameters
    datePrefrence := ctx.Param("datePrefrence")
    // get the user from the database
    userInterface, err := c.DB.Select(id)
    if err != nil {
        dr.InternalServerError(ctx, err)
        return
    }
    if userInterface == nil {
        dr.NotFound(ctx, fmt.Errorf("User with id %s not found", id))
        return
    }
    // unmarshal the ser from the database into a User object
    user := &models.User{}
    if err := surrealdb.Unmarshal(userInterface, user); err != nil {
        dr.InternalServerError(ctx, err)
        return
    }

    // update the user's date preference
    // and convert the date preference to an int
    user.DatePrefrence, err  = strconv.Atoi(datePrefrence)
    if err != nil {
        dr.BadRequest(ctx, err)
        return
    }

    // update the user in the database
    if _, err := c.DB.Update(user.Id, user); err != nil {
        dr.InternalServerError(ctx, err)
        return
    }
    dr.Data = user
    dr.OK(ctx)
}

// UpdateDatePrefrenceAuth handles the updating of the user 
// date preference.
// It validates the HMAC and calls the updateDatePrefrence function.
func (c *UserController) UpdateDatePrefrenceAuth(ctx *gin.Context) {
    dr := common.NewDetailedResponse(nil)
    if err := common.ValidateHmac(ctx); err != nil {
        dr.Unauthorized(ctx, err)
        return
    }

    c.updateDatePrefrence(ctx, dr)
}

// UpdateDatePrefrence handles the HTTP PUT request to update the user
// date preference.
// It calls the updateDatePrefrence method passing the context and a DetailedResponse object.
//
// WARNING: This method is only for development purposes.
func (c *UserController) UpdateDatePrefrence(ctx *gin.Context) {
    dr := common.NewDetailedResponse(nil)
    c.updateDatePrefrence(ctx, dr)
}

// CheckLogin handles if the user is logged in.
// This will later use a date to see if the user 
// has been logged in for too long. and force them to log in again.
func (c *UserController) checkLogin(ctx *gin.Context, dr *common.BoolResponse) {
    // get the user id from the request parameters
    authToken := ctx.Param("authToken")
    // get the user from the database
    userInterface, err := c.DB.Query(
        "SELECT * FROM User WHERE AuthToken = $authToken",
        map[string]interface{}{
            "authToken": authToken,
        })
    if err != nil {
        dr.InternalServerError(ctx, err)
        return
    }
    if userInterface == nil {
        dr.NotFound(ctx, fmt.Errorf("User with auth token %s not found", authToken))
        return
    }
    // unmarshal the ser from the database into a User object
    user := &models.User{}
    if err := surrealdb.Unmarshal(userInterface, user); err != nil {
        dr.InternalServerError(ctx, err)
        return
    }
    dr.OK(ctx)
}

// CheckLoginAuth handles the checking of the user login.
// It validates the HMAC and calls the checkLogin function.
func (c *UserController) CheckLoginAuth(ctx *gin.Context) {
    dr := common.NewBoolResponse()
    if err := common.ValidateHmac(ctx); err != nil {
        dr.Unauthorized(ctx, err)
        return
    }

    c.checkLogin(ctx, dr)
}

// CheckLogin handles the HTTP GET request to check if the user is logged in.
// It calls the checkLogin method passing the context and a DetailedResponse object.
//
// WARNING: This method is only for development purposes.
func (c *UserController) CheckLogin(ctx *gin.Context) {
    dr := common.NewBoolResponse()
    c.checkLogin(ctx, dr)
}
