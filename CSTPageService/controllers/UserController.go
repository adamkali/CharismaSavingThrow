package controllers

import (
    "github.com/gin-gonic/gin"
    common "github.com/adamkali/CharismaSavingThrow/CSTCommonLib"
    userService "github.com/adamkali/CharismaSavingThrow/CSTCommonLib/services"
    models "github.com/adamkali/CharismaSavingThrow/CSTCommonLib/models"
)

func Create(ctx *gin.Context) {
    dr := common.NewDetailedResponse(nil)
    userRequest := &models.UserRequest{}
    if err := ctx.Bind(userRequest); err != nil {
        dr.BadRequest(ctx, err)
        return
    }
    user, err := userService.NewUserService().Create(userRequest)
    if err != nil {
        dr.InternalServerError(ctx, err)
        return
    }
    dr.Data = user
    dr.OK(ctx)
}
