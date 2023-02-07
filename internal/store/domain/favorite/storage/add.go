package storage

import (
	"context"
	"github.com/eNViDAT0001/GolangAdventure/external/wrap_gorm"
	"github.com/eNViDAT0001/GolangAdventure/internal/store/entities"
	"gorm.io/gorm/clause"
)

func (f favoriteStorage) AddFavorite(ctx context.Context, userID uint, providerID uint) error {
	db := wrap_gorm.GetDB()
	subscribe := entities.Favorite{
		UserID:     userID,
		ProviderID: providerID,
	}

	err := db.Table(entities.Favorite{}.TableName()).Clauses(clause.OnConflict{DoNothing: true}).Create(&subscribe).Error
	if err != nil {
		return err
	}

	return nil
}
