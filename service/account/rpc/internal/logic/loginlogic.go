package logic

import (
	"context"
	"demo/common/errorx"
	"demo/service/account/model"
	"demo/service/account/rpc/internal/svc"
	"demo/service/account/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := l.svcCtx.UserModel.FindOneByName(l.ctx, in.Name)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errorx.NewDefaultError("not found user")
	default:
		return nil, err
	}
	if user.Password != in.Password {
		return nil, errorx.NewDefaultError("wrong password")
	}
	return &pb.LoginResponse{
		UserId: user.Id,
	}, nil
}
