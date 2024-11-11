// Copyright 2024 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package nodes

import (
	"github.com/tjbrains/CloudLogger/internal/configs"
	"github.com/tjbrains/CloudLogger/internal/remotelogs"
	"github.com/tjbrains/CloudLogger/pkg/rpc/pb"
	"github.com/tjbrains/CloudLogger/pkg/rpc/services"
	"github.com/tjbrains/TeaGo/types"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	rpcServer *grpc.Server
}

func NewServer() *Server {
	var options = []grpc.ServerOption{
		grpc.MaxRecvMsgSize(512 << 20),
		grpc.MaxSendMsgSize(512 << 20),
	}

	var rpcServer = grpc.NewServer(options...)
	pb.RegisterHttpAccessLogServiceServer(rpcServer, &services.HttpAccessLogService{})

	return &Server{
		rpcServer: rpcServer,
	}
}

func (this *Server) Listen() error {
	config, err := configs.LoadConfig()
	if err != nil {
		remotelogs.Error("SERVER", "load config failed: "+err.Error())
		remotelogs.Error("SERVER", "will use default config")
		config = configs.DefaultConfig()
	}

	remotelogs.Println("SERVER", "listening ':"+types.String(config.Port)+"' ...")

	listener, err := net.Listen("tcp", ":"+types.String(config.Port))
	if err != nil {
		return err
	}

	return this.rpcServer.Serve(listener)
}
