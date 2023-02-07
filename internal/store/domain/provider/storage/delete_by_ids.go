package storage

import (
	"context"
	"github.com/eNViDAT0001/GolangAdventure/external/wrap_gorm"
	"github.com/eNViDAT0001/GolangAdventure/internal/store/entities"
)

func (s providerStorage) DeleteProviderByIDs(ctx context.Context, providerID []uint) error {
	tableName := entities.Provider{}.TableName()
	db := wrap_gorm.GetDB()

	err := db.Table(tableName).
		Where("id IN ?", providerID).
		Delete(&entities.Provider{}).Error
	if err != nil {
		return err
	}
	return nil
}
