package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/x/errors"
	"go/types"
	"net/http"
	"wallet/api/server/internal/myMiddleware"

	"wallet/api/server/internal/config"
	"wallet/api/server/internal/handler"
	"wallet/api/server/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	xhttp "github.com/zeromicro/x/http"
)

var configFile = flag.String("f", "etc/server-api.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	server.Use(myMiddleware.PanicMiddleware)
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	// httpx.SetErrorHandler 仅在调用了 httpx.Error 处理响应时才有效。
	httpx.SetErrorHandler(func(err error) (int, any) {
		fmt.Println("SetErrorHandler:", err)
		switch e := err.(type) {
		case *errors.CodeMsg:
			return http.StatusOK, xhttp.BaseResponse[types.Nil]{
				Code: e.Code,
				Msg:  e.Msg,
			}
		default:
			return http.StatusInternalServerError, xhttp.BaseResponse[types.Nil]{
				Code: 50001,
				Msg:  e.Error(),
			}
		}
	})
	server.Start()
}
