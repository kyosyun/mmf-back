package repository

import "github.com/kyosyun/mmf-back/internal/domain/model"

type Repository interface {
	GetMurderMysteries() ([]model.MurderMystery, error)
	GetMurderMystery(id int) (model.MurderMystery, error)
	// UpdateMurderMysteryImage(id int, imageURL string) error
}
