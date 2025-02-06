package converter

import (
	domainModel "github.com/kyosyun/mmf-back/internal/domain/model"
	"github.com/kyosyun/mmf-back/internal/infra/postgres/model"
)

func ConvertMurderMysteryMasterToMurderMystery(murderMysteryMaster model.MurderMysteryMaster) domainModel.MurderMystery {
	return domainModel.MurderMystery{
		ID:              int(murderMysteryMaster.ID),
		JanCode:         murderMysteryMaster.JanCode,
		AsinCode:        murderMysteryMaster.AsinCode,
		Title:           murderMysteryMaster.Title,
		Description:     murderMysteryMaster.Description,
		Price:           int(murderMysteryMaster.Price),
		ReleaseDate:     murderMysteryMaster.ReleaseDate,
		Platform:        murderMysteryMaster.Platform,
		Genre:           murderMysteryMaster.Genre,
		PlayersMin:      int(murderMysteryMaster.PlayersMin),
		PlayersMax:      int(murderMysteryMaster.PlayersMax),
		AmazonLink:      murderMysteryMaster.AmazonLink,
		OfficialLink:    murderMysteryMaster.OfficialLink,
		ImageLink:       murderMysteryMaster.ImageLink,
		Language:        murderMysteryMaster.Language,
		Discount:        int(murderMysteryMaster.Discount),
		OriginalPrice:   int(murderMysteryMaster.OriginalPrice),
		PlayTime:        int(murderMysteryMaster.PlayTime),
		OnlineSupported: murderMysteryMaster.OnlineSupported,
		RequiresGm:      murderMysteryMaster.RequiresGm,
		SteamLink:       murderMysteryMaster.SteamLink,
	}
}

func ConvertMurderMysteryMastersToMurderMysteries(murderMysteryMasters []model.MurderMysteryMaster) []domainModel.MurderMystery {
	murderMysteries := make([]domainModel.MurderMystery, len(murderMysteryMasters))
	for i, murderMysteryMaster := range murderMysteryMasters {
		murderMysteries[i] = ConvertMurderMysteryMasterToMurderMystery(murderMysteryMaster)
	}
	return murderMysteries
}
