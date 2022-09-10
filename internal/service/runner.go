package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"strings"
	"time"
)

type Result struct {
	StdOut     string
	ResultCode int
	StdError   string
}

// run Executes a path with the desired runtime limit without performing any validation of the path
func run(ctx context.Context, path string, limit time.Duration, args string) (*Result, error) {
	if limit.Nanoseconds() > 0 {
		log.Debugf("processing custom runtime limit of %dms", limit)
		it, cf := context.WithTimeout(ctx, limit)
		ctx = it
		defer func() {
			cf()
		}()
	}

	log.Debugf("arguments: %s", args)
	stdout, err := exec.CommandContext(ctx, path, strings.Fields(args)...).Output()

	result := Result{
		StdOut:     strings.TrimSuffix(string(stdout), "\n"),
		ResultCode: 0,
	}

	if err != nil {
		if processError, ok := err.(*exec.ExitError); ok {
			result.StdError = string(processError.Stderr)
			result.ResultCode = processError.ExitCode()
		}
	}

	return &result, nil

}
