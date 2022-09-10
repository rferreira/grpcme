package main

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpcme/internal/pb"
	"net"
	"os"
	"os/exec"
)

func main() {
	verbose := flag.Bool("v", false, "verbose mode enabled, spews out a lot more logs to the console")
	healthCheckArgument := flag.String("h", "--help", "healthcheck argument to be used")
	address := flag.String("a", "127.0.0.1:8089", "set the listening host:port combination")
	noHealthcheck := flag.Bool("noHealthCheck", false, "skips health checking the executable")

	flag.Parse()
	if flag.NArg() == 0 {
		println("Error: You must pass in the full path to an executable to wrap")
		print(fmt.Sprintf("For example: .%s  /bin/date\n", os.Args[0]))
		flag.Usage()
		os.Exit(1)
	}
	// configure logging
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	if *verbose {
		log.SetLevel(log.DebugLevel)
		log.Debugf("Running in verbose mode...")
	}

	executable := flag.Arg(0)
	log.Infof("wrapping executable %s\n", executable)
	log.Debugf("using health check argument: %s", *healthCheckArgument)
	path, err := exec.LookPath(executable)
	if err != nil {
		log.Fatal(err)
	}
	if !*noHealthcheck {
		out, err := exec.Command(path, *healthCheckArgument).Output()
		if err != nil {
			log.Fatal(err)
		}
		if *verbose {
			log.Debugf("preflight output: %s", string(out))
		}
	} else {
		log.Warnf("skipping health check")
	}

	log.Printf("pre-flight check passed, starting GRPC service...")
	var opts []grpc.ServerOption
	lis, err := net.Listen("tcp", *address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on %s", *address)
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterGrpcMeServer(grpcServer, pb.NewGrpcMeServer(path))
	reflection.Register(grpcServer)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
