package postgres

import (
	domainModel "github.com/kyosyun/mmf-back/internal/domain/model"
	"github.com/kyosyun/mmf-back/internal/infra/postgres/converter"
	"github.com/kyosyun/mmf-back/internal/infra/postgres/model"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) GetMurderMysteries() ([]domainModel.MurderMystery, error) {
	var murderMysteryMasters []model.MurderMysteryMaster

	if err := r.db.Find(&murderMysteryMasters).Error; err != nil {
		return []domainModel.MurderMystery{}, err
	}

	return converter.ConvertMurderMysteryMastersToMurderMysteries(murderMysteryMasters), nil
}

func (r *PostgresRepository) GetMurderMystery(id int) (domainModel.MurderMystery, error) {
	var mm model.MurderMysteryMaster
	if err := r.db.First(&mm, id).Error; err != nil {
		return domainModel.MurderMystery{}, err
	}
	return converter.ConvertMurderMysteryMasterToMurderMystery(mm), nil
}
