package main

import (
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"myshorturl/api/internal/config"
	"myshorturl/api/internal/handler"
	"myshorturl/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/shorturl-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	logx.DisableStat()
	server.Start()
}
