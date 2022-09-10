package main

import (
	"context"
	"flag"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpcme/internal/pb"
)

func main() {
	verbose := flag.Bool("v", false, "verbose mode enabled, spews out a lot more logs to the console")
	server := flag.String("a", "127.0.0.1:8089", "set the server host:port combination")
	args := flag.String("d", "", "argument to be passed to the server")
	//secure := flag.Bool("s", false, "should we use TLS or not to connect to the client")

	flag.Parse()

	// configure logging
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	if *verbose {
		log.SetLevel(log.DebugLevel)
		log.Debugf("Running in verbose mode...")
	}

	conn, err := grpc.Dial(*server, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewGrpcMeClient(conn)

	exec, err := client.Exec(context.Background(), &pb.ExecRequest{
		Args: args,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("Execution with result code: %d", exec.ResultCode)
	log.Infof("Stdout: %s", exec.StdOut)
	log.Infof("Stderr: %s", exec.StdError)
}
