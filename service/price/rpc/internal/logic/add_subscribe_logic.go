package logic

import (
	"context"
	"demo/common/errorx"
	"demo/service/price/model/sql"
	"github.com/go-sql-driver/mysql"

	"demo/service/price/rpc/internal/svc"
	"demo/service/price/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSubscribeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSubscribeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSubscribeLogic {
	return &AddSubscribeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddSubscribeLogic) AddSubscribe(in *pb.AddSubscribeRequest) (*pb.Empty, error) {
	_, err := l.svcCtx.SubscribeModel.Insert(l.ctx, &sql.Subscribe{
		UserId: in.Subscribe.UserId,
		ItemId: in.Subscribe.ItemId,
		ShopId: in.Subscribe.ShopId,
	})
	switch e := err.(type) {
	case nil:
	case *mysql.MySQLError:
		if e.Number == 1062 {
			return nil, errorx.ErrDuplicateSubscribe
		}
	default:
		return nil, err
	}
	return &pb.Empty{}, nil
}
