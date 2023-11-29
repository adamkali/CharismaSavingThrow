package controller

import (
	"net/http"

	models "github.com/adamkali/CharismaSavingThrow/CSTCommonLib/models"
	userService "github.com/adamkali/CharismaSavingThrow/CSTCommonLib/services"
	"github.com/gin-gonic/gin"
)


func Create(ctx *gin.Context) {
    userRequest := &models.UserRequest{}
    if err := ctx.Bind(userRequest); err != nil {
        return
    }
    //user, err := userService.NewUserService().Create(userRequest)
    //if err != nil {
    //    return
    //}
    
}

type createUserFormPageData struct {
    UsernameInput InputComponent
    PasswordInput InputComponent
    ConfirmPasswordInput InputComponent
    EmailInput InputComponent
    SubmitButton SubmitButtonComponent
    BioInput InputComponent
    DatePreferenceInput InputComponent
    LoginButton ButtonComponent
}


func CheckLoggedIn(ctx *gin.Context) {
    authToken := ctx.Param("authToken")
    loggedIn, err := userService.NewUserService().CheckLoggedIn(authToken)
    if err != nil {
        return
    }
    if !loggedIn {
        // send the CreatUserFormPage
        createUserFormPage := &createUserFormPageData{
            UsernameInput: InputComponent{
                Name: "username",
                Type: "text",
                Value: "",
                Placeholder: "Username",
                Disabled: false,
            },
            PasswordInput: InputComponent{
                Name: "password",
                Type: "password",
                Value: "",
                Placeholder: "Password",
                Disabled: false,
            },
            ConfirmPasswordInput: InputComponent{
                Name: "confirmPassword",
                Type: "password",
                Value: "",
                Placeholder: "Confirm Password",
                Disabled: false,
            },
            EmailInput: InputComponent{
                Name: "email",
                Type: "email",
                Value: "",
                Placeholder: "Email",
                Disabled: false,
            },
            SubmitButton: SubmitButtonComponent{
                Name: "submit",
                Text: "Create Account",
                Icon: "static/imgs/.svg",
            },
            BioInput: InputComponent{
                Name: "bio",
                Type: "text",
                Value: "",
                Placeholder: "Bio",
            },
            DatePreferenceInput: InputComponent{
                Name: "datePreference",
                Type: "number",
                Value: "1",
                Placeholder: "Date Preference",
                Disabled: true,
            },
        }
        ctx.HTML(http.StatusOK, "user/createForm", createUserFormPage)
    }
    // send the user page
    ctx.HTML(http.StatusOK, "user/userPage", nil)
}
