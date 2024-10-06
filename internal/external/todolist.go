package external

import (
	"context"
)

type TodoList interface {
	List(ctx context.Context) ([]string, error)
	Add(ctx context.Context, text string) (string, error)
	Remove(ctx context.Context, id string) (string, error)
}

//go:generate gowrap gen -i TodoList -g -t ../templates/opentracing -o todolist/generated/todo_with_spans.go
//go:generate gowrap gen -i TodoList -g -t ../templates/prometheus -o todolist/generated/todo_with_metrics.go
//go:generate gowrap gen -i TodoList -g -t ../templates/logrus -o todolist/generated/todo_with_logs.go
//go:generate gowrap gen -i TodoList -g -t ../templates/timeout -o todolist/generated/todo_with_timeout.go
//go:generate gowrap gen -i TodoList -g -t ../templates/ratelimit -o todolist/generated/todo_with_ratelimit.go
//go:generate gowrap gen -i TodoList -g -t ../templates/retry -o todolist/generated/todo_with_retry.go
//go:generate gowrap gen -i TodoList -g -t ../templates/circuitbreaker -o todolist/generated/todo_with_breaker.go
