package services

import (
	"context"
	"fmt"
	"os"

	cstcommonlib "github.com/adamkali/CharismaSavingThrow/CSTCommonLib"
	"github.com/adamkali/CharismaSavingThrow/CSTCommonLib/models"
	"github.com/carlmjohnson/requests"
)


type DatePreferenceService struct {
    endpoint string
}

func NewDatePreferenceService() *DatePreferenceService {
    return &DatePreferenceService{
        endpoint: os.Getenv("CST_DATA_ENDPOINT"),
    }
}

func (dps *DatePreferenceService) GetByNumber(dpNumber int) (*models.DatePreference, error) {
    endpoint := dps.endpoint + "/api/auth/datePreference/" + string(rune(dpNumber))
    HmacAuthHeader, err := cstcommonlib.ConstructHmacAuthHeader(endpoint, "GET")
    if err != nil {
        return nil, err
    }

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
    dp := &models.DatePreference{}
    dp = dr.Data.(*models.DatePreference)
    return dp, nil
}

func (dps *DatePreferenceService) GetAll() ([]*models.DatePreference, error) {
    endpoint := dps.endpoint + "/api/auth/datePreference/"
    HmacAuthHeader, err := cstcommonlib.ConstructHmacAuthHeader(endpoint, "GET")
    if err != nil {
        return nil, err
    }

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
    dprefs := []*models.DatePreference{}
    dprefs = dr.Data.([]*models.DatePreference)
    return dprefs, nil
}
