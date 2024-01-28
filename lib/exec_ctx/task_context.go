package core_concept

import (
	"context"
	"time"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/rs/zerolog"
)

type TaskContext struct {
	execName string

	Ctx        context.Context
	Err        chan error
	CancelTask func()
	LastError  error
}

func (ctx *TaskContext) GetExecutionName() string {
	return ctx.execName
}
func (ctx *TaskContext) GetDurationStat() time.Duration {
	return time.Minute
}

func (ctx *TaskContext) GetCtx() context.Context {
	return ctx.Ctx
}

func (ctx *TaskContext) Cancel() {
	ctx.CancelTask()
}

func (ctx *TaskContext) SetErrorCustom(err error, handler func(event *zerolog.Event) *zerolog.Event) {
	ctx.LastError = err
	pdc_common.ReportErrorCustom(err, handler)
	ctx.Err <- err
}

func (ctx *TaskContext) SetError(err error) {
	pdc_common.ReportError(err)
	ctx.LastError = err
	ctx.Err <- err
}

func NewTaskContext(ctx context.Context, execName string) *TaskContext {
	ctx, cancel := context.WithCancel(ctx)
	errchan := make(chan error, 10)
	return &TaskContext{
		execName: execName,

		Ctx: ctx,
		Err: errchan,
		CancelTask: func() {
			cancel()
			close(errchan)
		},
	}
}
