package controller

import (
	"net/http"

	models "github.com/adamkali/CharismaSavingThrow/CSTCommonLib/models"
	"github.com/adamkali/CharismaSavingThrow/CSTCommonLib/services"
	"github.com/gin-gonic/gin"
)

type datePreferenceSelectorData struct {
    DatePreferences []models.DatePreference
}

func DatePreferenceSelector(ctx *gin.Context) {
	var datePreferences []*models.DatePreference
    datePreferences, err := services.NewDatePreferenceService().GetAll()
    if err != nil {
        // TODO: handle error
        return
    }
    
    datePreferencesSelectorData := datePreferenceSelectorData{
        DatePreferences: make([]models.DatePreference, len(datePreferences)),
    }
    for i, datePreference := range datePreferences {
        datePreferencesSelectorData.DatePreferences[i] = *datePreference
    }

	ctx.HTML(
        http.StatusOK,
        "datePreferences/date-preference-selector.html",
        datePreferencesSelectorData,
    )
}
