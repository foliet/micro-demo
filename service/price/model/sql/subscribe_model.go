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
		FindAllByUserId(ctx context.Context, userId int64) ([]*Subscribe, error)
	}

	customSubscribeModel struct {
		*defaultSubscribeModel
	}
)

// NewSubscribeModel returns a model for the database table.
func NewSubscribeModel(conn sqlx.SqlConn) SubscribeModel {
	return &customSubscribeModel{
		defaultSubscribeModel: newSubscribeModel(conn),
	}
}

func (m *customSubscribeModel) FindAllByUserId(ctx context.Context, userId int64) ([]*Subscribe, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ?", subscribeRows, m.table)
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
