package storage

import (
	"context"
	"github.com/eNViDAT0001/GolangAdventure/external/paging"
	"github.com/eNViDAT0001/GolangAdventure/external/paging/paging_query"
	"github.com/eNViDAT0001/GolangAdventure/external/wrap_gorm"
	"github.com/eNViDAT0001/GolangAdventure/internal/store/entities"
)

func (s providerStorage) CountListProviderByUserID(ctx context.Context, userID uint, filter paging.ParamsInput) (total int64, err error) {
	db := wrap_gorm.GetDB()
	query := db.Model(entities.Provider{}).Where("user_id = ?", userID)
	paging_query.SetCountListPagingQuery(&filter, entities.Provider{}.TableName(), query)
	err = query.Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}
