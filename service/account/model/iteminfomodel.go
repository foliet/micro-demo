package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ItemInfoModel = (*customItemInfoModel)(nil)

type (
	// ItemInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customItemInfoModel.
	ItemInfoModel interface {
		itemInfoModel
	}

	customItemInfoModel struct {
		*defaultItemInfoModel
	}
)

// NewItemInfoModel returns a model for the database table.
func NewItemInfoModel(conn sqlx.SqlConn) ItemInfoModel {
	return &customItemInfoModel{
		defaultItemInfoModel: newItemInfoModel(conn),
	}
}
