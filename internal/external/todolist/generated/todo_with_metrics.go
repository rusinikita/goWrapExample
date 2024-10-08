// Code generated by gowrap. DO NOT EDIT.
// template: ../../../templates/prometheus
// gowrap: http://github.com/hexdigest/gowrap

package generated

import (
	"context"
	_sourceExternal "goWrapExample/internal/external"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// TodoListWithPrometheus implements _sourceExternal.TodoList interface with all methods wrapped
// with Prometheus metrics
type TodoListWithPrometheus struct {
	base         _sourceExternal.TodoList
	instanceName string
}

var todolistDurationSummaryVec = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "todolist_duration_seconds",
		Help:       "todolist runtime duration and result",
		MaxAge:     time.Minute,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"instance_name", "method", "result"})

// NewTodoListWithPrometheus returns an instance of the _sourceExternal.TodoList decorated with prometheus summary metric
func NewTodoListWithPrometheus(base _sourceExternal.TodoList, instanceName string) TodoListWithPrometheus {
	return TodoListWithPrometheus{
		base:         base,
		instanceName: instanceName,
	}
}

// Add implements _sourceExternal.TodoList
func (_d TodoListWithPrometheus) Add(ctx context.Context, text string) (s1 string, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		todolistDurationSummaryVec.WithLabelValues(_d.instanceName, "Add", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.Add(ctx, text)
}

// List implements _sourceExternal.TodoList
func (_d TodoListWithPrometheus) List(ctx context.Context) (sa1 []string, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		todolistDurationSummaryVec.WithLabelValues(_d.instanceName, "List", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.List(ctx)
}

// Remove implements _sourceExternal.TodoList
func (_d TodoListWithPrometheus) Remove(ctx context.Context, id string) (s1 string, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		todolistDurationSummaryVec.WithLabelValues(_d.instanceName, "Remove", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.Remove(ctx, id)
}
