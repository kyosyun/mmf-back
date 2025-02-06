package model

import (
	"fmt"
	"time"
)

type MurderMystery struct {
	ID              int
	JanCode         string
	AsinCode        string
	Title           string
	Description     string
	Price           int
	ReleaseDate     time.Time
	Platform        string
	Genre           string
	PlayersMin      int
	PlayersMax      int
	AmazonLink      string
	OfficialLink    string
	ImageLink       string
	Language        string
	Discount        int
	OriginalPrice   int
	PlayTime        int
	OnlineSupported bool
	RequiresGm      bool
	SteamLink       string
}

func (m *MurderMystery) GetAmazonAffiliateLink(tag string) string {
	return fmt.Sprintf("https://www.amazon.co.jp/gp/product/%s/?tag=%s", m.AsinCode, tag)
}
