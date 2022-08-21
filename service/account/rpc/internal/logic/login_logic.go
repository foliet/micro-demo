package logic

import (
	"context"
	"demo/common/errorx"
	"demo/service/account/model/sql"
	"demo/service/account/rpc/internal/svc"
	"demo/service/account/rpc/pb"
	"golang.org/x/crypto/bcrypt"

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

func (l *LoginLogic) Login(in *pb.LoginRequest) (*pb.UserId, error) {
	user, err := l.svcCtx.UserModel.FindOneByName(l.ctx, in.Name)
	switch err {
	case nil:
	case sql.ErrNotFound:
		return nil, errorx.ErrUsernameNotFound
	default:
		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)) != nil {
		return nil, errorx.ErrWrongPassword
	}
	return &pb.UserId{
		Id: user.Id,
	}, nil
}
