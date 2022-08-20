package logic

import (
	"context"
	"demo/common/errorx"
	"demo/service/account/model/sql"
	"demo/service/account/rpc/internal/svc"
	"demo/service/account/rpc/pb"
	"github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/codes"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterRequest) (*pb.UserId, error) {
	if in.Code != "1234" {
		return nil, errorx.NewDefaultError("wrong verification code")
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	result, err := l.svcCtx.UserModel.Insert(l.ctx, &sql.User{
		Name:     in.Name,
		Password: string(hashPassword),
	})
	switch e := err.(type) {
	case *mysql.MySQLError:
		if e.Number == 1062 {
			return nil, errorx.NewCodeError(codes.Code(e.Number), "name had been registered")
		}
	default:
		return nil, err
	}
	id, err := result.LastInsertId()
	return &pb.UserId{Id: id}, nil
}
