package ddos

import (
	"context"
	"ddos-grpc/internal/services/ddos"

	dds_v1 "github.com/jantttez/ddos-proto/gen/go/ddos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ddosApi struct {
	dds_v1.UnimplementedDDosServer
	D *ddos.DDos
}

const (
	InitSuccess   = "DDos init Success run toggle option to start"
	StopToggle    = "DDos stop toggle success"
	StartToggle   = "DDos start toggle success"
	StartOption   = "start"
	StopOption    = "stop"
	emptyValue    = ""
	emptyIntValue = 0
)

func RegisterGrpcServer(gRpc *grpc.Server) {
	dds_v1.RegisterDDosServer(gRpc, &ddosApi{})
}

func (d *ddosApi) DDosInit(
	ctx context.Context,
	req *dds_v1.DDosInitRequest,
) (*dds_v1.DDosResponse, error) {
	url := req.GetDdosUrl()
	amount := req.GetRequestCount()
	e := make(chan bool)

	if url == emptyValue || amount == emptyIntValue {
		return nil, status.Error(codes.InvalidArgument, "invalid payload")
	}

	if d.D == nil {
		d.D = &ddos.DDos{
			ReqAmount: amount,
			Url:       url,
			Exit:      &e,
		}
	}

	d.D.ReqAmount = amount
	d.D.Url = url
	d.D.Exit = &e

	return &dds_v1.DDosResponse{
		Status: InitSuccess,
	}, nil
}

func (d *ddosApi) DDosToggle(
	ctx context.Context,
	req *dds_v1.DDosToggleRequest,
) (*dds_v1.DDosResponse, error) {
	option := req.GetOption()
	if option == emptyValue {
		return nil, status.Error(codes.InvalidArgument, "option is empty")
	}

	if option == StartOption {
		d.D.Run()
	} else if option == StopOption {
		d.D.Stop()
	}

	if option == StartOption {
		return &dds_v1.DDosResponse{
			Status: StartToggle,
		}, nil

	} else {
		return &dds_v1.DDosResponse{
			Status: StopToggle,
		}, nil
	}
}
