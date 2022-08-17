package logic

import (
	"context"
	"demo/common/errorx"
	"demo/service/account/model"
	"demo/service/account/rpc/pb/account"

	"demo/service/account/rpc/internal/svc"

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

func (l *LoginLogic) Login(in *account.LoginRequest) (*account.LoginResponse, error) {
	user, err := l.svcCtx.UserModel.FindOneByName(l.ctx, in.Name)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errorx.NewDefaultError("用户名不存在")
	default:
		return nil, err
	}
	if user.Password != in.Password {
		return nil, errorx.NewDefaultError("密码错误")
	}
	return &account.LoginResponse{
		UserId: user.Id,
	}, nil
}
