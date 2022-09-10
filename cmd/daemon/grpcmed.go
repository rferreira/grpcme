package main

import (
	"github.com/pelletier/go-toml/v2"
	log "github.com/sirupsen/logrus"
	"grpcme/internal/server"
	"io/ioutil"
	"os"
)

var usage = `Usage:
	./grpcmed CONFIG-FILE.toml
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

	var config = server.Config{}
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

	err, srv := server.New(config)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(srv.String())

	//executable := flag.Arg(0)
	//log.Printf("pre-flight check passed, starting GRPC service...")
	//var opts []grpc.ServerOption
	//lis, err := net.Listen("tcp", *address)
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//log.Printf("Listening on %s", *address)
	//grpcServer := grpc.NewServer(opts...)
	//pb.RegisterGrpcMeServer(grpcServer, pb.NewGrpcMeServer(path))
	//reflection.Register(grpcServer)
	//err = grpcServer.Serve(lis)
	//if err != nil {
	//	log.Fatal(err)
	//}
}
