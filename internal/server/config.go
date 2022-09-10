package server

import "time"

type Config struct {
	Verbose        bool
	DefaultTimeout durationWrapper                 `toml:"default_timeout"`
	Executables    map[string]ExecutableDefinition `toml:"executable"`
}

type ExecutableDefinition struct {
	Path                string
	HealthCheckArgument string `toml:"healthcheck_argument"`
}

type durationWrapper struct {
	Value time.Duration
}

func (w *durationWrapper) UnmarshalText(text []byte) error {
	val, err := time.ParseDuration(string(text))
	if err != nil {
		return err
	}
	w.Value = val
	return nil
}
