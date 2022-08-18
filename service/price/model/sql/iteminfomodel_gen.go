// Code generated by goctl. DO NOT EDIT!

package sql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	itemInfoFieldNames          = builder.RawFieldNames(&ItemInfo{})
	itemInfoRows                = strings.Join(itemInfoFieldNames, ",")
	itemInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(itemInfoFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	itemInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(itemInfoFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	itemInfoModel interface {
		Insert(ctx context.Context, data *ItemInfo) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ItemInfo, error)
		Update(ctx context.Context, data *ItemInfo) error
		Delete(ctx context.Context, id int64) error
	}

	defaultItemInfoModel struct {
		conn  sqlx.SqlConn
		table string
	}

	ItemInfo struct {
		Id       int64     `db:"id"`
		CreateAt time.Time `db:"create_at"`
		ItemId   int64     `db:"item_id"`
		Price    float64   `db:"price"`
	}
)

func newItemInfoModel(conn sqlx.SqlConn) *defaultItemInfoModel {
	return &defaultItemInfoModel{
		conn:  conn,
		table: "`item_info`",
	}
}

func (m *defaultItemInfoModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultItemInfoModel) FindOne(ctx context.Context, id int64) (*ItemInfo, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", itemInfoRows, m.table)
	var resp ItemInfo
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
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
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, itemInfoRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.ItemId, data.Price)
	return ret, err
}

func (m *defaultItemInfoModel) Update(ctx context.Context, data *ItemInfo) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, itemInfoRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.ItemId, data.Price, data.Id)
	return err
}

func (m *defaultItemInfoModel) tableName() string {
	return m.table
}
