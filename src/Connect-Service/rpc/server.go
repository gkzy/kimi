package rpc

import (
	"github.com/gkzy/gow/lib/logy"
	"github.com/gkzy/gow/lib/rpc"
	"github.com/gkzy/kimi/common/constname"
	"github.com/gkzy/kimi/common/repository/pb"
	"google.golang.org/grpc"
)

// InitGRPCServer start rpc server
func InitGRPCServer() {
	g, err := rpc.NewServer(constname.ConnectServerRPCPort)
	if err != nil {
		panic(err)
	}
	handler(g.Server)
	g.Run()

	logy.Infof("[GRPC] Listening and serving TCP on %v", g.Port)
}

func handler(s *grpc.Server) {
	pb.RegisterConnMessageServer(s,&ConnMessageGRPCService{})
}
