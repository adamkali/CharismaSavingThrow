package controller

import (
	"net/http"

	models "github.com/adamkali/CharismaSavingThrow/CSTCommonLib/models"
	"github.com/adamkali/CharismaSavingThrow/CSTCommonLib/services"
	"github.com/gin-gonic/gin"
)

func DatePreferenceSelector(ctx *gin.Context) {
	var datePreferences []*models.DatePreference
    datePreferences, err := services.NewDatePreferenceService().GetAll()
    if err != nil {
        // TODO: handle error
        return
    }

	ctx.HTML(http.StatusOK, "date-preference-selector.html", gin.H{
		"datePreferences": datePreferences,
	})
}
