package server

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os/exec"
)

type Server struct {
	Config *Config
}

func New(config Config) (error, *Server) {
	log.Infof("Parsing exposed executables...")
	found := 0
	for k, v := range config.Executables {
		log.Printf("[%s] -> %s", k, v.Path)
		validationErr := validate(v, config.Verbose)
		if validationErr != nil {
			return fmt.Errorf("validated failed for executable: %s, %w", v.Path, validationErr), nil
		}
		found++
	}

	if found == 0 {
		return errors.New("no executables defined in config file"), nil
	}
	return nil, &Server{
		Config: &config,
	}
}

func (receiver Server) String() string {
	return fmt.Sprintf("Server version %d", 0)
}

func validate(it ExecutableDefinition, verbose bool) error {
	log.Infof("wrapping executable %s", it.Path)
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
		return fmt.Errorf("exit code:%d stderr:\n%s", wrappedError.ExitCode(), wrappedError.Stderr)
	}
	if verbose {
		log.Debugf("validation output: %s", string(out))
	}
	return nil
}
