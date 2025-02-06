package handler

import (
	"context"

	"github.com/kyosyun/mmf-back/api"
	"github.com/kyosyun/mmf-back/internal/domain/model"
	"github.com/kyosyun/mmf-back/internal/domain/repository"
	"github.com/kyosyun/mmf-back/internal/env"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

var _ api.StrictServerInterface = &App{}

type App struct {
	Repository repository.Repository
}

func NewApp(repository repository.Repository) *App {
	return &App{Repository: repository}
}

func (a *App) GetMurderMysteries(ctx context.Context, req api.GetMurderMysteriesRequestObject) (api.GetMurderMysteriesResponseObject, error) {

	mm, err := a.Repository.GetMurderMysteries()
	if err != nil {
		return api.GetMurderMysteries400Response{}, err
	}

	return api.GetMurderMysteries200JSONResponse(convertMurderMysteriesToApi(mm)), nil
}

func (a *App) GetMurderMysteriesId(ctx context.Context, req api.GetMurderMysteriesIdRequestObject) (api.GetMurderMysteriesIdResponseObject, error) {

	return api.GetMurderMysteriesId200JSONResponse{}, nil
}

func convertMurderMysteryToApi(mm *model.MurderMystery) *api.MurderMystery {
	releaseDate := openapi_types.Date{Time: mm.ReleaseDate}
	amazonLink := mm.GetAmazonAffiliateLink(env.GetEnvironment().AmazonConfig.PartnerTag)
	return &api.MurderMystery{
		Id:           mm.ID,
		Title:        mm.Title,
		ReleaseDate:  releaseDate,
		ImageLink:    &mm.ImageLink,
		AmazonLink:   &amazonLink,
		OfficialLink: &mm.OfficialLink,
	}
}

func convertMurderMysteriesToApi(mms []model.MurderMystery) []api.MurderMystery {
	result := make([]api.MurderMystery, len(mms))
	for i, mm := range mms {
		result[i] = *convertMurderMysteryToApi(&mm)
	}
	return result
}
