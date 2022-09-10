package service

import (
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"time"
)

type ValidationError struct {
	Path     string
	StdError string
	ExitCode int
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("validation failed for executable: %s, exit code: %d, stderr:\n%s", v.Path, v.ExitCode, v.StdError)
}

type Service struct {
	Config *Config
}

func New(config *Config) (*Service, error) {
	log.Infof("Parsing exposed executables...")
	found := 0
	for k, v := range config.Executables {
		log.Printf("[%s] -> %s", k, v.Path)
		validationErr := validate(v, config.Verbose)
		if validationErr != nil {
			return nil, validationErr
		}
		found++
	}
	if found == 0 {
		return nil, errors.New("no executables defined in config file")
	}
	return &Service{
		Config: config,
	}, nil
}

func (receiver *Service) String() string {
	return fmt.Sprintf("Server version %d", 0)
}

func validate(it ExecutableDefinition, verbose bool) error {
	path, err := exec.LookPath(it.Path)
	if err != nil {
		return err
	}
	log.Debugf("using health check argument: %s", it.HealthCheckArgument)

	var out []byte
	var execError error

	if it.HealthCheckArgument == "" {
		out, execError = exec.Command(path).Output()
	} else {
		out, execError = exec.Command(path, it.HealthCheckArgument).Output()
	}
	if execError != nil {
		wrappedError := execError.(*exec.ExitError)
		return ValidationError{
			Path:     path,
			StdError: string(wrappedError.Stderr),
			ExitCode: wrappedError.ExitCode(),
		}
	}
	if verbose {
		log.Debugf("validation output: %s", string(out))
	}
	return nil
}

func (receiver *Service) Handle(ctx context.Context, id string, timeOut time.Duration, args string) (*Result, error) {
	var path string

	// should we honor the default timeout?
	if receiver.Config.DefaultTimeout.Value != 0 && timeOut == 0 {
		log.Debugf("using default time for call without an explicit timeout")
		timeOut = receiver.Config.DefaultTimeout.Value
	}

	if def, found := receiver.Config.Executables[id]; found {
		path = def.Path
	} else {
		return nil, errors.New("unknown executable with id: " + id)
	}

	return run(ctx, path, timeOut, args)
}
