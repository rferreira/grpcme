package service

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	type args struct {
		ctx          context.Context
		path         string
		limit        time.Duration
		optionalArgs string
	}
	tests := []struct {
		name    string
		args    args
		want    *Result
		wantErr bool
	}{
		{name: "date", args: args{ctx: context.TODO(), path: "/bin/date", limit: time.Hour * 1, optionalArgs: "+%Y"},
			want: &Result{
				StdOut:     fmt.Sprintf("%d", time.Now().Year()),
				ResultCode: 0,
			}, wantErr: false},
		{name: "date", args: args{ctx: context.TODO(), path: "/bin/date", limit: time.Hour * 1, optionalArgs: "+%Y"},
			want: &Result{
				StdOut:     fmt.Sprintf("%d", time.Now().Year()),
				ResultCode: 0,
			}, wantErr: false},
		{name: "date", args: args{ctx: context.TODO(), path: "/bin/date", limit: time.Hour * 1, optionalArgs: "--help"},
			want: &Result{
				ResultCode: 1,
				StdError:   "/bin/date: illegal option -- -\nusage: date [-jnRu] [-r seconds|file] [-v[+|-]val[ymwdHMS]]\n            [-I[date | hours | minutes | seconds]]\n            [-f fmt date | [[[mm]dd]HH]MM[[cc]yy][.ss]] [+format]\n",
			}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := run(tt.args.ctx, tt.args.path, tt.args.limit, tt.args.optionalArgs)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run() got = %v, want %v", got, tt.want)
			}
		})
	}
}
