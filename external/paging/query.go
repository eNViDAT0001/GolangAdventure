package paging

import (
	"fmt"
	"github.com/eNViDAT0001/Backend/external/request"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPagingParams(cc *gin.Context, filter EntityFilter) (paginator GetListInput, err error) {
	if err = cc.BindQuery(&paginator); err != nil {
		return paginator, err
	}

	search := cc.QueryArray("search[]")
	fields := cc.QueryArray("fields[]")
	sort := cc.QueryArray("sorts[]")

	paginator.Filter = NewFilterBuilder().
		WithSearch(search).
		WithFields(fields).
		WithSorts(sort).
		Build()

	inValidField, val := ValidateFilter(paginator.Filter, filter)
	if len(inValidField) > 0 {
		return paginator, request.NewBadRequestError(inValidField, val, "invalid key and value")
	}
	return paginator, err
}

func SetCountListPagingQuery(input *GetListInput, tableName string, query *gorm.DB) {

	if input.Type == CURSOR_PAGING {
		query = query.Where(fmt.Sprintf("%s.id > ?", tableName), input.Current())
	}

	if input.Filter.GetFields() != nil {
		for k, v := range *input.Filter.GetFields() {
			query = query.Where(fmt.Sprintf("`%s`.`%s` = ?", tableName, k), v)
		}
	}

	if input.Filter.GetSearch() != nil {
		for k, v := range *input.Filter.GetSearch() {
			query = query.Where(fmt.Sprintf("`%s`.`%s` LIKE ?", tableName, k), "%"+v+"%")
		}
	}

	if input.Filter.GetSort() != nil {
		for k, v := range *input.Filter.GetSort() {
			sort := "ASC"
			if v == "DESC" {
				sort = v
			}
			query = query.Order(fmt.Sprintf(`%s.%s %s`, tableName, k, sort))
		}
	}
}
func SetPagingQuery(input *GetListInput, tableName string, query *gorm.DB) {

	query = query.Limit(input.PerPage())
	if input.Type == CURSOR_PAGING {
		query = query.Where(fmt.Sprintf("%s.id > ?", tableName), input.Current())
	} else {
		offset := Offset(input.Current(), input.PerPage())
		query = query.Offset(offset)
	}

	if input.Filter.GetFields() != nil {
		for k, v := range *input.Filter.GetFields() {
			query = query.Where(fmt.Sprintf("`%s`.`%s` = ?", tableName, k), v)
		}
	}

	if input.Filter.GetSearch() != nil {
		for k, v := range *input.Filter.GetSearch() {
			query = query.Where(fmt.Sprintf("`%s`.`%s` LIKE ?", tableName, k), "%"+v+"%")
		}
	}

	if input.Filter.GetSort() != nil {
		for k, v := range *input.Filter.GetSort() {
			sort := "ASC"
			if v == "DESC" {
				sort = v
			}
			query = query.Order(fmt.Sprintf(`%s.%s %s`, tableName, k, sort))
		}
	}
}
