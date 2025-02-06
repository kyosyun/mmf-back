package handler

import (
	"context"

	"github.com/kyosyun/mmf-back/api"
)

func (a *App) GetHealth(ctx context.Context, req api.GetHealthRequestObject) (api.GetHealthResponseObject, error) {
	return api.GetHealth200JSONResponse{
		Message: "ok",
	}, nil
}
