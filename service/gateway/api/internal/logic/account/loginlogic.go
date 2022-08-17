package account

import (
	"context"
	"demo/service/account/rpc/pb/account"
	"demo/service/gateway/api/internal/svc"
	"demo/service/gateway/api/internal/types"
	"github.com/golang-jwt/jwt/v4"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	resp = new(types.LoginResponse)
	rpcResp, err := l.svcCtx.AccountRpc.Login(l.ctx, &account.LoginRequest{
		Name:     req.Name,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	resp.AccessToken, resp.AccessExpire, err = l.getJwtToken(rpcResp.GetUserId())
	return
}

func (l *LoginLogic) getJwtToken(userId int64) (string, int64, error) {
	auth := l.svcCtx.Config.Auth
	iat := time.Now().Unix()
	exp := iat + auth.AccessExpire
	claims := make(jwt.MapClaims)
	claims["exp"] = exp
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	tokenString, err := token.SignedString([]byte(auth.AccessSecret))
	if err != nil {
		return "", 0, err
	}
	return tokenString, exp, nil
}
