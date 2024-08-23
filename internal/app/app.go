package app

import (
	"ddos-grpc/internal/app/grpc"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	gRPCApp *grpc.GrpcApp
}

func NewApp() *App {
	gRPCApp := grpc.NewGrpcApp(":7777")

	return &App{
		gRPCApp: gRPCApp,
	}
}

func (a *App) Run() {
	log.Print("grpc server is running")

	a.gRPCApp.MustRun()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)

	<-exit
	log.Print("grpc server stoping...")

	a.gRPCApp.Stop()

	log.Print("grpc server stopped")
}
