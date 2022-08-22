package sql

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ItemInfoModel = (*customItemInfoModel)(nil)

type (
	// ItemInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customItemInfoModel.
	ItemInfoModel interface {
		itemInfoModel
		FindAllByUserIdAndItemId(ctx context.Context, page int64, userId int64, itemId int64) ([]*ItemInfo, error)
	}

	customItemInfoModel struct {
		*defaultItemInfoModel
		pageSize int64
	}
)

// NewItemInfoModel returns a model for the database table.
func NewItemInfoModel(conn sqlx.SqlConn) ItemInfoModel {
	return &customItemInfoModel{
		defaultItemInfoModel: newItemInfoModel(conn),
		pageSize:             20,
	}
}

func (m *customItemInfoModel) FindAllByUserIdAndItemId(ctx context.Context, page int64, userId int64, itemId int64) ([]*ItemInfo, error) {
	query := fmt.Sprintf("select %s from (select `item_id` from `subscribe` where `user_id` = ?) as subscribe natural join item_info where `item_id` = ? limit %d offset %d", itemInfoRows, m.pageSize, m.pageSize*page)
	var resp []*ItemInfo
	err := m.conn.QueryRowsCtx(ctx, &resp, query, userId, itemId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
