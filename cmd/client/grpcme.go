package main

import (
	"context"
	"fmt"
	"github.com/rferreira/grpcme/internal/pb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"strings"
)

var usage = `Usage:
	./grpcme t./grpcme TARGET ID [ARGS]
for example: 
	./grpcme dns:localhost:8089 date

For help with the correct GRPC url please see: https://github.com/grpc/grpc/blob/master/doc/naming.md
`

func main() {
	if len(os.Args) == 1 {
		print(usage)
		os.Exit(1)
	}

	// configure logging
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	if os.Getenv("GRPCMEDEBUG") != "" {
		log.SetLevel(log.DebugLevel)
		log.Debugf("Running in verbose mode...")
	}

	target := os.Args[1]
	id := os.Args[2]

	var args string
	if len(os.Args) > 3 {
		args = strings.Join(os.Args[3:len(os.Args)], " ")
	}
	// https://github.com/grpc/grpc/blob/master/doc/naming.md
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewGrpcMeClient(conn)

	exec, err := client.Exec(context.Background(), &pb.ExecRequest{
		Id:   id,
		Args: &args,
	})
	if err != nil {
		log.Fatal(err)
	}

	if exec.ResultCode != 0 {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("Execution with result code: %d\n", exec.ResultCode))
		_, _ = os.Stderr.WriteString(fmt.Sprintf("Stdout:\n%s", exec.StdOut))
		_, _ = os.Stderr.WriteString(fmt.Sprintf("Stderr:\n%s", exec.StdError))
		os.Exit(1)
	}

	fmt.Println(exec.StdOut)
}
