package sql

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SubscribeModel = (*customSubscribeModel)(nil)

type (
	// SubscribeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSubscribeModel.
	SubscribeModel interface {
		subscribeModel
		FindAllByUserId(ctx context.Context, page int64, userId int64) ([]*Subscribe, error)
		FindAllWithDistinctItemId(ctx context.Context) ([]*Subscribe, error)
	}

	customSubscribeModel struct {
		*defaultSubscribeModel
		pageSize int64
	}
)

// NewSubscribeModel returns a model for the database table.
func NewSubscribeModel(conn sqlx.SqlConn) SubscribeModel {
	return &customSubscribeModel{
		defaultSubscribeModel: newSubscribeModel(conn),
		pageSize:              20,
	}
}

func (m *customSubscribeModel) FindAllByUserId(ctx context.Context, page int64, userId int64) ([]*Subscribe, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? limit %d offset %d", subscribeRows, m.table, m.pageSize, m.pageSize*page)
	var resp []*Subscribe
	err := m.conn.QueryRowsCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customSubscribeModel) FindAllWithDistinctItemId(ctx context.Context) ([]*Subscribe, error) {
	query := fmt.Sprintf("select %s from %s", subscribeRows, m.table)
	var resp []*Subscribe
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
