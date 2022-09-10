package service

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
	"io/fs"
	"os/exec"
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	type args struct {
		config *Config
	}

	tests := []struct {
		name string
		args args
		want error
	}{
		{name: "should require executables defined", args: args{
			&Config{},
		}, want: errors.New("no executables defined in config file")},
		{name: "should startup if executable validation passes", args: args{
			&Config{Executables: map[string]ExecutableDefinition{
				"date": {Path: "/bin/date"},
			}},
		}, want: nil},
		{name: "should error if executable validation fails", args: args{
			&Config{Executables: map[string]ExecutableDefinition{
				"foo": {Path: "/bin/foo"},
			}},
		}, want: &exec.Error{
			Name: "/bin/foo",
			Err: &fs.PathError{
				Op:   "stat",
				Path: "/bin/foo",
				Err:  unix.ENOENT,
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got := New(tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_HandleDefaultTimeOut(t *testing.T) {
	srv, err := New(&Config{
		DefaultTimeout: durationWrapper{1 * time.Second},
		Executables: map[string]ExecutableDefinition{
			// sadly we need to pass an argument to sleep so it doesn't error out during validation
			"sleep": {Path: "/bin/sleep", HealthCheckArgument: "0.01"},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	result, err := srv.Handle(context.Background(), "sleep", 0, "2")
	if err != nil {
		log.Fatal(err)
	}
	if result.ResultCode != -1 {
		t.Errorf("execution did not timeout")
	}
}
