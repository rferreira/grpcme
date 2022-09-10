package pb

import (
	"context"
	log "github.com/sirupsen/logrus"
	"grpcme/internal/runner"
	"time"
)

func NewGrpcMeServer(path string) GrpcMeServer {
	return &DefaultGrpcMeServer{
		path: path,
	}
}

type DefaultGrpcMeServer struct {
	path string
}

func (it DefaultGrpcMeServer) Exec(ctx context.Context, request *ExecRequest) (*ExecResponse, error) {
	startTime := time.Now()
	defer func() {
		log.Infof("request completed in %s", time.Now().Sub(startTime))
	}()

	var duration time.Duration
	if request.Limit.IsValid() {
		duration = request.GetLimit().AsDuration()
	}

	var args string

	if request.Args != nil {
		args = *request.Args
	}

	result, err := runner.Run(ctx, it.path, duration, args)
	if err != nil {
		return nil, err
	}

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
