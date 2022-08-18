package logic

import (
	"context"
	"demo/common/errorx"
	"demo/service/price/model/sql"
	"demo/service/price/rpc/internal/svc"
	"demo/service/price/rpc/pb"
	"github.com/go-sql-driver/mysql"
	codes "google.golang.org/grpc/codes"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubscribeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubscribeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubscribeLogic {
	return &SubscribeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubscribeLogic) Subscribe(in *pb.SubscribeRequest) (*pb.Empty, error) {
	_, err := l.svcCtx.SubscribeModel.Insert(l.ctx, &sql.Subscribe{
		UserId: in.UserId,
		ItemId: in.ItemId,
		ShopId: in.ShopId,
	})
	switch e := err.(type) {
	case *mysql.MySQLError:
		if e.Number == 1062 {
			return nil, errorx.NewCodeError(codes.Code(e.Number), "had been subscribed")
		}
	default:
		return nil, err
	}
	return &pb.Empty{}, nil
}
