package account

import (
	"context"
	"demo/service/account/rpc/account"
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
	auth := l.svcCtx.Config.Auth
	resp.AccessToken, resp.AccessExpire, err = getJwtToken(rpcResp.GetId(), auth.AccessSecret, auth.AccessExpire)
	return
}

func getJwtToken(userId int64, secret string, expire int64) (string, int64, error) {
	iat := time.Now().Unix()
	exp := iat + expire
	claims := make(jwt.MapClaims)
	claims["exp"] = exp
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", 0, err
	}
	return tokenString, exp, nil
}
