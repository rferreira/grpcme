package main

import (
	"github.com/pelletier/go-toml/v2"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpcme/internal/pb"
	"grpcme/internal/service"
	"io/ioutil"
	"net"
	"os"
)

var usage = `Usage:
	./grpcmed config.toml
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

	filename := os.Args[1]
	log.Printf("parsing %s", filename)

	var config = service.Config{}
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	err = toml.Unmarshal(contents, &config)
	if err != nil {
		log.Errorf("error processing config file %s", filename)
		log.Fatal(err)
	}

	if config.Verbose {
		log.SetLevel(log.DebugLevel)
		log.Debugf("Running in verbose mode...")
	}

	srv, err := service.New(&config)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(srv.String())

	var opts []grpc.ServerOption
	lis, err := net.Listen("tcp", config.Listen)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on %s", config.Listen)
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterGrpcMeServer(grpcServer, pb.NewGrpcMeServer(srv))
	reflection.Register(grpcServer)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
