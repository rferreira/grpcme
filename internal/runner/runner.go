package runner

import (
	"context"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"time"
)

type Result struct {
	StdOut     string
	ResultCode int
	StdError   string
}

// Run Executes a path with the desired runtime limit without performing any validation of the path
func Run(ctx context.Context, path string, limit time.Duration, optionalArgs string) (*Result, error) {
	if limit.Nanoseconds() > 0 {
		log.Debugf("processing custom runtime limit of %dms", limit)
		it, cf := context.WithTimeout(ctx, limit)
		ctx = it
		defer func() {
			cf()
		}()
	}

	var args []string

	if len(optionalArgs) > 0 {
		args = append(args, optionalArgs)
	}
	log.Debugf("arguments: %s", args)
	stdout, err := exec.CommandContext(ctx, path, args...).Output()

	result := Result{
		StdOut:     string(stdout),
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
