package local

import (
	"configcenter/src/common/blog"
	"context"

	"go.mongodb.org/mongo-driver/event"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
	"go.opentelemetry.io/otel/trace"
)

// SQL日志记录
func newCommandMonitor() *event.CommandMonitor {
	return &event.CommandMonitor{
		Started: func(ctx context.Context, startedEvent *event.CommandStartedEvent) {
			blog.Infof("mongo reqId:%d start on db:%s cmd:%s sql:%+v", startedEvent.RequestID, startedEvent.DatabaseName,
				startedEvent.CommandName, startedEvent.Command)
		},
		Succeeded: func(ctx context.Context, succeededEvent *event.CommandSucceededEvent) {
			blog.Infof("mongo reqId:%d exec cmd:%s success duration %d ns", succeededEvent.RequestID,
				succeededEvent.CommandName, succeededEvent.DurationNanos)
		},
		Failed: func(ctx context.Context, failedEvent *event.CommandFailedEvent) {
			blog.Infof("mongo reqId:%d exec cmd:%s failed duration %d ns", failedEvent.RequestID,
				failedEvent.CommandName, failedEvent.DurationNanos)
		},
	}
}

// 链路追踪
func newOtelMonitor(tp trace.TracerProvider) *event.CommandMonitor {
	return otelmongo.NewMonitor(otelmongo.WithTracerProvider(tp), otelmongo.WithCommandAttributeDisabled(false))
}

func combineMonitors(monitors ...*event.CommandMonitor) *event.CommandMonitor {
	return &event.CommandMonitor{
		Started: func(ctx context.Context, startedEvent *event.CommandStartedEvent) {
			for _, m := range monitors {
				m.Started(ctx, startedEvent)
			}
		},
		Succeeded: func(ctx context.Context, succeededEvent *event.CommandSucceededEvent) {
			for _, m := range monitors {
				m.Succeeded(ctx, succeededEvent)
			}
		},
		Failed: func(ctx context.Context, failedEvent *event.CommandFailedEvent) {
			for _, m := range monitors {
				m.Failed(ctx, failedEvent)
			}
		},
	}
}
