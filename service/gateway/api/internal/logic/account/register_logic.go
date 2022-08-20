package account

import (
	"context"
	"demo/service/account/rpc/pb"

	"demo/service/gateway/api/internal/svc"
	"demo/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	resp = new(types.RegisterResponse)
	rpcResp, err := l.svcCtx.AccountRpc.Register(l.ctx, &pb.RegisterRequest{
		Name:     req.Name,
		Password: req.Password,
		Code:     req.Code,
	})
	if err != nil {
		return nil, err
	}
	auth := l.svcCtx.Config.Auth
	resp.AccessToken, resp.AccessExpire, err = getJwtToken(rpcResp.GetId(), auth.AccessSecret, auth.AccessExpire)
	return
}
