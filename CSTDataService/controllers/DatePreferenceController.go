package controllers

import (
	"fmt"

	common "github.com/adamkali/CharismaSavingThrow/CSTCommonLib"
	models "github.com/adamkali/CharismaSavingThrow/CSTCommonLib/models"
	"github.com/gin-gonic/gin"
	"github.com/surrealdb/surrealdb.go"
)

type DatePreferenceController struct {
	DB *surrealdb.DB
}

func NewDatePreferenceController(db *surrealdb.DB, engine *gin.Engine) *DatePreferenceController {
     return &DatePreferenceController{DB: db}
}

func (date *DatePreferenceController) getDatePreference(ctx *gin.Context, dr *common.DetailedResponse) {
    dpNumber := ctx.Param("dpNumber")
    DatePreference := &models.DatePreference{}

    dp, err := date.DB.Query(
        "SELECT * FROM date_preferences WHERE number = $dpNumber",
        map[string]interface{}{
            "dpNumber": dpNumber,
        },
    )
    if err != nil {
        dr.InternalServerError(ctx, fmt.Errorf("An error occurred while querying the database: %s", err))
        return
    }
    if dp == nil {
        dr.NotFound(ctx, fmt.Errorf("Date Preference with numberID %s not found", dpNumber))
        return
    }
    if err := surrealdb.Unmarshal(dp, DatePreference); err != nil {
        dr.InternalServerError(ctx, fmt.Errorf("An error occurred while unmarshalling the date preference: %s", err))
        return
    }
    dr.Data = *DatePreference
    dr.OK(ctx)

}

func (date *DatePreferenceController) getAllDatePreferences(c *gin.Context, dr *common.DetailedResponse) {
    datePreferences := []*models.DatePreference{}
    nilDatePreference := &models.DatePreference{}
    dps, err := date.DB.Select(nilDatePreference.GetTableName())
    if err != nil {
        dr.InternalServerError(c, err)
        return
    }
    if dps == nil {
        dr.NotFound(c, fmt.Errorf("No date preferences found"))
        return
    }
    if err := surrealdb.Unmarshal(dps, &datePreferences); err != nil {
        dr.InternalServerError(c, err)
        return
    }
    dr.Data = datePreferences
    dr.OK(c)
}

func (date *DatePreferenceController) GetDatePreferenceAuth(ctx *gin.Context) {
    dr := common.NewDetailedResponse(nil)
    if err := common.ValidateHmac(ctx); err != nil {
        dr.Unauthorized(ctx, err)
        return
    }
    date.getDatePreference(ctx, dr)
}

func (date *DatePreferenceController) GetDatePreference(ctx *gin.Context) {
    dr := common.NewDetailedResponse(nil)
    date.getDatePreference(ctx, dr)
}
    
func (date *DatePreferenceController) GetAllDatePreferencesAuth(ctx *gin.Context) {
    dr := common.NewDetailedResponse(nil)
    if err := common.ValidateHmac(ctx); err != nil {
        dr.Unauthorized(ctx, err)
        return
    }
    date.getAllDatePreferences(ctx, dr)
}

func (date *DatePreferenceController) GetAllDatePreferences(ctx *gin.Context) {
    dr := common.NewDetailedResponse(nil)
    date.getAllDatePreferences(ctx, dr)
}
