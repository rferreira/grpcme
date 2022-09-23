package pb

import (
	"context"
	"github.com/rferreira/grpcme/internal/service"
	log "github.com/sirupsen/logrus"
	"time"
)

func NewExecServer(it *service.Service) ExecServer {
	return &DefaultExecServer{
		service: it,
	}
}

type DefaultExecServer struct {
	service *service.Service
}

func (it *DefaultExecServer) mustEmbedUnimplementedExecServer() {
	//TODO implement me
	panic("implement me")
}

func (it *DefaultExecServer) Exec(ctx context.Context, request *ExecRequest) (*ExecResponse, error) {
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
