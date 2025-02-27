// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// MurderMystery defines model for MurderMystery.
type MurderMystery struct {
	// AmazonLink Amazon purchase link
	AmazonLink *string `json:"amazon_link,omitempty"`

	// Id Unique identifier for the game
	Id int `json:"id"`

	// ImageLink URL of the game's image
	ImageLink *string `json:"image_link,omitempty"`

	// OfficialLink Official website link
	OfficialLink *string `json:"official_link,omitempty"`

	// ReleaseDate Release date of the game
	ReleaseDate openapi_types.Date `json:"release_date"`

	// Title Title of the game
	Title string `json:"title"`
}

// MurderMysteryDetail defines model for MurderMysteryDetail.
type MurderMysteryDetail struct {
	AgeRecommendation *string             `json:"age_recommendation,omitempty"`
	AmazonLink        *string             `json:"amazon_link,omitempty"`
	BookwalkerLink    *string             `json:"bookwalker_link,omitempty"`
	Description       *string             `json:"description,omitempty"`
	Discount          *float32            `json:"discount,omitempty"`
	DlsiteLink        *string             `json:"dlsite_link,omitempty"`
	EpicgamesLink     *string             `json:"epicgames_link,omitempty"`
	FanzaGamesLink    *string             `json:"fanza_games_link,omitempty"`
	Genre             *string             `json:"genre,omitempty"`
	Id                *int                `json:"id,omitempty"`
	ImageLink         *string             `json:"image_link,omitempty"`
	Language          *string             `json:"language,omitempty"`
	OfficialLink      *string             `json:"official_link,omitempty"`
	OnlineSupported   *bool               `json:"online_supported,omitempty"`
	OriginalPrice     *float32            `json:"original_price,omitempty"`
	Platform          *string             `json:"platform,omitempty"`
	PlayTime          *int                `json:"play_time,omitempty"`
	PlayersMax        *int                `json:"players_max,omitempty"`
	PlayersMin        *int                `json:"players_min,omitempty"`
	Price             *float32            `json:"price,omitempty"`
	RakutenBooksLink  *string             `json:"rakuten_books_link,omitempty"`
	ReleaseDate       *openapi_types.Date `json:"release_date,omitempty"`
	Replayability     *string             `json:"replayability,omitempty"`
	RequiresGm        *bool               `json:"requires_gm,omitempty"`
	SteamLink         *string             `json:"steam_link,omitempty"`
	SurugayaLink      *string             `json:"surugaya_link,omitempty"`
	Title             *string             `json:"title,omitempty"`
}

// GetMurderMysteriesParams defines parameters for GetMurderMysteries.
type GetMurderMysteriesParams struct {
	// YearMonth The year and month for which to retrieve games (YYYY-MM format)
	YearMonth *string `form:"yearMonth,omitempty" json:"yearMonth,omitempty"`
}
