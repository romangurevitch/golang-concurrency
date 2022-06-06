package logger

import (
	"context"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestTraceDuration(t *testing.T) {
	traceLogger := logrus.New()
	traceLogger.SetFormatter(&logrus.TextFormatter{})
	traceLogger.SetLevel(logrus.TraceLevel)
	Init(traceLogger)

	hook := test.NewLocal(log)

	type args struct {
		ctx   context.Context
		start time.Time
		msg   string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "simple1", args: args{ctx: context.Background(), start: time.Time{}, msg: "test1"}},
		{name: "simple2", args: args{ctx: context.Background(), start: time.Time{}.Add(1 * time.Second), msg: "test2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TraceDuration(tt.args.ctx, tt.args.start, tt.args.msg)
			assert.Equal(t, hook.LastEntry().Message, tt.args.msg)
			assert.Equal(t, hook.LastEntry().Level, logrus.TraceLevel)
		})
	}
}
