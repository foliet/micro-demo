// Code generated by goctl. DO NOT EDIT!

package sql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userSubscribeFieldNames          = builder.RawFieldNames(&UserSubscribe{})
	userSubscribeRows                = strings.Join(userSubscribeFieldNames, ",")
	userSubscribeRowsExpectAutoSet   = strings.Join(stringx.Remove(userSubscribeFieldNames, "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	userSubscribeRowsWithPlaceHolder = strings.Join(stringx.Remove(userSubscribeFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheDemoUserSubscribeIdPrefix = "cache:demo:userSubscribe:id:"
)

type (
	userSubscribeModel interface {
		Insert(ctx context.Context, data *UserSubscribe) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserSubscribe, error)
		Update(ctx context.Context, data *UserSubscribe) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserSubscribeModel struct {
		sqlc.CachedConn
		table string
	}

	UserSubscribe struct {
		UserId int64 `db:"user_id"`
		ItemId int64 `db:"item_id"`
		Id     int64 `db:"id"`
	}
)

func newUserSubscribeModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserSubscribeModel {
	return &defaultUserSubscribeModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user_subscribe`",
	}
}

func (m *defaultUserSubscribeModel) Delete(ctx context.Context, id int64) error {
	demoUserSubscribeIdKey := fmt.Sprintf("%s%v", cacheDemoUserSubscribeIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, demoUserSubscribeIdKey)
	return err
}

func (m *defaultUserSubscribeModel) FindOne(ctx context.Context, id int64) (*UserSubscribe, error) {
	demoUserSubscribeIdKey := fmt.Sprintf("%s%v", cacheDemoUserSubscribeIdPrefix, id)
	var resp UserSubscribe
	err := m.QueryRowCtx(ctx, &resp, demoUserSubscribeIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userSubscribeRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserSubscribeModel) Insert(ctx context.Context, data *UserSubscribe) (sql.Result, error) {
	demoUserSubscribeIdKey := fmt.Sprintf("%s%v", cacheDemoUserSubscribeIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, userSubscribeRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.ItemId, data.Id)
	}, demoUserSubscribeIdKey)
	return ret, err
}

func (m *defaultUserSubscribeModel) Update(ctx context.Context, data *UserSubscribe) error {
	demoUserSubscribeIdKey := fmt.Sprintf("%s%v", cacheDemoUserSubscribeIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userSubscribeRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.ItemId, data.Id)
	}, demoUserSubscribeIdKey)
	return err
}

func (m *defaultUserSubscribeModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheDemoUserSubscribeIdPrefix, primary)
}

func (m *defaultUserSubscribeModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userSubscribeRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserSubscribeModel) tableName() string {
	return m.table
}