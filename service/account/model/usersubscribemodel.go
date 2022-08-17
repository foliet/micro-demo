package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserSubscribeModel = (*customUserSubscribeModel)(nil)

type (
	// UserSubscribeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserSubscribeModel.
	UserSubscribeModel interface {
		userSubscribeModel
	}

	customUserSubscribeModel struct {
		*defaultUserSubscribeModel
	}
)

// NewUserSubscribeModel returns a model for the database table.
func NewUserSubscribeModel(conn sqlx.SqlConn) UserSubscribeModel {
	return &customUserSubscribeModel{
		defaultUserSubscribeModel: newUserSubscribeModel(conn),
	}
}
