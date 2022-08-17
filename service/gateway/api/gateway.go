package main

import (
	"demo/common/errorx"
	"demo/service/gateway/api/internal/config"
	"demo/service/gateway/api/internal/handler"
	"demo/service/gateway/api/internal/svc"
	"flag"
	"fmt"
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
		statusError, ok := status.FromError(err)
		if !ok {
			err = errorx.NewDefaultError(err.Error())
		}
		statusError = status.Convert(err)
		return http.StatusOK, &errorx.ErrorResponse{
			Code: statusError.Code(),
			Msg:  statusError.Message(),
		}
	})
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
