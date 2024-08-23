package grpc

import (
	"ddos-grpc/internal/grpc/ddos"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type GrpcApp struct {
	gRPCApp *grpc.Server
	port    string
}

func NewGrpcApp(p string) *GrpcApp {
	gRPCSrv := grpc.NewServer()

	ddos.RegisterGrpcServer(gRPCSrv)

	return &GrpcApp{
		gRPCApp: gRPCSrv,
		port:    p,
	}
}
func (ga *GrpcApp) MustRun() {
	if err := ga.Run(); err != nil {
		panic(err)
	}
}

// сначала создаещб листенер tcp пакетов и порта и затем добавляешь его на прослушку в grpc сервер
func (ga *GrpcApp) Run() error {

	l, err := net.Listen("tcp", ga.port)
	if err != nil {
		return fmt.Errorf("error with listen tcp port, and error looks like: %s", err.Error())
	}

	log.Printf("grpc server is running on addres: %s", l.Addr().String())

	if err := ga.gRPCApp.Serve(l); err != nil {
		return fmt.Errorf("error with listemn grpc server %s", err)
	}

	return nil
}

func (ga *GrpcApp) Stop() {

	ga.gRPCApp.GracefulStop()
}
