package pb

import (
	"context"
	log "github.com/sirupsen/logrus"
	"grpcme/internal/service"
	"time"
)

func NewGrpcMeServer(it *service.Service) GrpcMeServer {
	return &DefaultGrpcMeServer{
		service: it,
	}
}

type DefaultGrpcMeServer struct {
	service *service.Service
}

func (it DefaultGrpcMeServer) Exec(ctx context.Context, request *ExecRequest) (*ExecResponse, error) {
	startTime := time.Now()

	var duration time.Duration
	if request.Limit.IsValid() {
		duration = request.GetLimit().AsDuration()
	}

	var args string

	if request.Args != nil {
		args = *request.Args
	}

	result, err := it.service.Handle(ctx, request.Id, duration, args)
	if err != nil {
		log.Errorf("Error: %s", err)
		return nil, err
	}

	defer func() {
		log.Infof("REQ: [%s] completed in %s with exit code: %d", request.GetId(), time.Now().Sub(startTime), result.ResultCode)
	}()

	return &ExecResponse{
		StdOut:     result.StdOut,
		ResultCode: int32(result.ResultCode),
		StdError:   result.StdError,
	}, err
}

func (DefaultGrpcMeServer) mustEmbedUnimplementedGrpcMeServer() {
	//TODO implement me
	panic("implement me")
}
