package main

import (
	"github.com/cloudwego/kitex/server"
	"github.com/kouleen/common/bootstrap"
	"github.com/kouleen/common/middleware"
	"github.com/kouleen/idl/kitex_gen/rpc"
	"github.com/kouleen/idl/kitex_gen/stream/streamservice"
)

func main() {
	bootstrap.Run(rpc.STREAM_RPC_SERVER, func(option ...server.Option) server.Server {
		return streamservice.NewServer(new(StreamServiceImpl), option...)
	}, bootstrap.WithServerMiddleware(middleware.RpcServerMiddleware))
}
