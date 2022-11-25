/**
 *******************************************51cto********************************************
 * Copyright (c)  www.51cto.com
 * Created by go-sensitive.
 * User: shijl@51cto.com
 * Date: 2021/08/26
 * Time: 11:18
 ********************************************************************************************
 */

package main

import (
	"flag"
	"fmt"

	"go-sensitive/api/internal/config"
	"go-sensitive/api/internal/handler"
	"go-sensitive/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/sensitive-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
