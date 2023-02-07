package storage

import (
	"context"
	"github.com/eNViDAT0001/GolangAdventure/external/wrap_gorm"
	"github.com/eNViDAT0001/GolangAdventure/internal/store/entities"
)

func (b bannerStorage) DeleteBannerByIDs(ctx context.Context, bannerIDs []uint) error {
	tableName := entities.BannerDetail{}.TableName()
	db := wrap_gorm.GetDB()

	err := db.Table(tableName).Where("id IN ?", bannerIDs).Delete(&entities.Banner{}).Error

	if err != nil {
		return err
	}

	return nil
}
