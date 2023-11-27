package services

import (
	"context"
	"fmt"
	"os"

	cstcommonlib "github.com/adamkali/CharismaSavingThrow/CSTCommonLib"
	"github.com/adamkali/CharismaSavingThrow/CSTCommonLib/models"
	"github.com/carlmjohnson/requests"
)

// allow for other serveces to use the user service
// by provideng a user service struct to access the
// user service api
type UserService struct {
    endpoint string
}

// NewUserService creates a new user service struct
// and will use the environment variables to set the 
// endpoint
func NewUserService() *UserService {
    return &UserService{
        endpoint: os.Getenv("CST_USER_ENDPOINT"),
    }
}

// Create will create a new user in the user service
// and return the user object
//
// userRequest: the user object to Create
func (u *UserService) Create(userRequest *models.UserRequest) (*models.User, error) {
    endpoint := u.endpoint + "/api/auth/user/create"
    HmacAuthHeader, err := cstcommonlib.ConstructHmacAuthHeader(endpoint, "POST")
    if err != nil {
        return nil, err
    }
    //construct the request
    //	headers := map[string]interface{}{
	//	"Signature":      hmacSignature,
	//	"Timestamp":      timestamp,
	//	"Content-Type":   "application/json",
	//}

    var dr cstcommonlib.DetailedResponse
    err = requests.
        URL(endpoint).
        Header("Signature", HmacAuthHeader["Signature"]).
        Header("Timestamp", HmacAuthHeader["Timestamp"]).
        Header("Content-Type", HmacAuthHeader["Content-Type"]).
        BodyJSON(userRequest).
        ToJSON(dr).
        Fetch(context.Background())
    if err != nil {
        return nil, err
    }
    // check if the request was successful
    if !dr.Success {
        return nil, fmt.Errorf(dr.Message)
    }
    // get the user object from the response 
    user := &models.User{}
    user = dr.Data.(*models.User)
    return user, nil
}

// Get will get a user from the user service and return
// the user object
func (u *UserService) Get(id string) (*models.User, error) {
    endpoint := u.endpoint + "/api/auth/user/" + id
    HmacAuthHeader, err := cstcommonlib.ConstructHmacAuthHeader(endpoint, "GET")
    if err != nil {
        return nil, err
    }
    //construct the request
    //	headers := map[string]interface{}{
    //	"Signature":      hmacSignature,
    //	"Timestamp":      timestamp,
    //	"Content-Type":   "application/json",
    //}

    var dr cstcommonlib.DetailedResponse
    err = requests.
        URL(endpoint).
        Header("Signature", HmacAuthHeader["Signature"]).
        Header("Timestamp", HmacAuthHeader["Timestamp"]).
        Header("Content-Type", HmacAuthHeader["Content-Type"]).
        ToJSON(dr).
        Fetch(context.Background())
    if err != nil {
        return nil, err
    }
    // check if the request was successful
    if !dr.Success {
        return nil, fmt.Errorf(dr.Message)
    }
    // get the user object from the response 
    user := &models.User{}
    user = dr.Data.(*models.User)
    return user, nil
}

