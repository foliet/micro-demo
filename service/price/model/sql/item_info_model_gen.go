// Code generated by goctl. DO NOT EDIT!

package sql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	itemInfoFieldNames          = builder.RawFieldNames(&ItemInfo{})
	itemInfoRows                = strings.Join(itemInfoFieldNames, ",")
	itemInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(itemInfoFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	itemInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(itemInfoFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheDemoItemInfoIdPrefix = "cache:demo:itemInfo:id:"
)

type (
	itemInfoModel interface {
		Insert(ctx context.Context, data *ItemInfo) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ItemInfo, error)
		Update(ctx context.Context, data *ItemInfo) error
		Delete(ctx context.Context, id int64) error
	}

	defaultItemInfoModel struct {
		sqlc.CachedConn
		table string
	}

	ItemInfo struct {
		Id       int64     `db:"id"`
		CreateAt time.Time `db:"create_at"`
		ItemId   int64     `db:"item_id"`
		Price    float64   `db:"price"`
	}
)

func newItemInfoModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultItemInfoModel {
	return &defaultItemInfoModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`item_info`",
	}
}

func (m *defaultItemInfoModel) Delete(ctx context.Context, id int64) error {
	demoItemInfoIdKey := fmt.Sprintf("%s%v", cacheDemoItemInfoIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, demoItemInfoIdKey)
	return err
}

func (m *defaultItemInfoModel) FindOne(ctx context.Context, id int64) (*ItemInfo, error) {
	demoItemInfoIdKey := fmt.Sprintf("%s%v", cacheDemoItemInfoIdPrefix, id)
	var resp ItemInfo
	err := m.QueryRowCtx(ctx, &resp, demoItemInfoIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", itemInfoRows, m.table)
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

func (m *defaultItemInfoModel) Insert(ctx context.Context, data *ItemInfo) (sql.Result, error) {
	demoItemInfoIdKey := fmt.Sprintf("%s%v", cacheDemoItemInfoIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, itemInfoRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ItemId, data.Price)
	}, demoItemInfoIdKey)
	return ret, err
}

func (m *defaultItemInfoModel) Update(ctx context.Context, data *ItemInfo) error {
	demoItemInfoIdKey := fmt.Sprintf("%s%v", cacheDemoItemInfoIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, itemInfoRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.ItemId, data.Price, data.Id)
	}, demoItemInfoIdKey)
	return err
}

func (m *defaultItemInfoModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheDemoItemInfoIdPrefix, primary)
}

func (m *defaultItemInfoModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", itemInfoRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultItemInfoModel) tableName() string {
	return m.table
}
