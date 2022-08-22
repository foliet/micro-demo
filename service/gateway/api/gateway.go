package main

import (
	"demo/service/gateway/api/internal/config"
	"demo/service/gateway/api/internal/handler"
	"demo/service/gateway/api/internal/svc"
	"demo/service/gateway/api/internal/types"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/gateway.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		logx.Error(err.Error())
		statusError, ok := status.FromError(err)
		if !ok || statusError.Code() < 1000 {
			return http.StatusInternalServerError, nil
		}
		statusError = status.Convert(err)
		return http.StatusOK, &types.CodeResponse{
			Code: int64(statusError.Code()),
			Msg:  statusError.Message(),
		}
	})
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
