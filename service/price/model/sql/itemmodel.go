package sql

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ItemModel = (*customItemModel)(nil)

type (
	// ItemModel is an interface to be customized, add more methods here,
	// and implement the added methods in customItemModel.
	ItemModel interface {
		itemModel
	}

	customItemModel struct {
		*defaultItemModel
	}
)

// NewItemModel returns a model for the database table.
func NewItemModel(conn sqlx.SqlConn, c cache.CacheConf) ItemModel {
	return &customItemModel{
		defaultItemModel: newItemModel(conn, c),
	}
}
