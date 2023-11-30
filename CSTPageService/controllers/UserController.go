package controller

import (
	"net/http"

	models "github.com/adamkali/CharismaSavingThrow/CSTCommonLib/models"
	userService "github.com/adamkali/CharismaSavingThrow/CSTCommonLib/services"
	"github.com/gin-gonic/gin"
)


func UserPage(ctx *gin.Context) {
    userId := ctx.Param("userId")
    user, err := userService.NewUserService().Get(userId)
    if err != nil {
        // TODO: handle error
        return
    }
    ctx.HTML(http.StatusOK, "user/index.html", user)
}

func Create(ctx *gin.Context) {
    userRequest := &models.UserRequest{}
    if err := ctx.Bind(userRequest); err != nil {
        return
    }
    user, err := userService.NewUserService().Create(userRequest)
    if err != nil {
        // TODO: handle error
        return
    }
    // redirect to the user page with the user's id
    ctx.Redirect(http.StatusFound, "/user/" + user.Id)
}

type createUserFormPageData struct {
    UsernameInput InputComponent
    PasswordInput InputComponent
    ConfirmPasswordInput InputComponent
    EmailInput InputComponent
    SubmitButton SubmitButtonComponent
    BioInput InputComponent
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
                Label: "Username",
                Type: "text",
                Value: "",
                Disabled: false,
            },
            PasswordInput: InputComponent{
                Label: "Password",
                Name: "password",
                Type: "password",
                Value: "",
                Disabled: false,
            },
            ConfirmPasswordInput: InputComponent{
                Label: "Confirm Password",
                Name: "confirmPassword",
                Type: "password",
                Value: "",
                Disabled: false,
            },
            EmailInput: InputComponent{
                Label: "Email",
                Name: "email",
                Type: "email",
                Value: "",
                Disabled: false,
            },
            SubmitButton: SubmitButtonComponent{
                Text: "Create User",
                SvgName: "user",
            },
            BioInput: InputComponent{
                Name: "bio",
                Type: "text",
                Value: "",
                Disabled: false,
                Label: "Bio",
            },
        }
        ctx.HTML(http.StatusOK, "user/createForm", createUserFormPage)
    }
    // send the user page
    ctx.HTML(http.StatusOK, "user/userPage", nil)
}
