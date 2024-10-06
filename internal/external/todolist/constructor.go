package todolist

import (
	"database/sql"
	"time"

	"goWrapExample/internal/external"
	"goWrapExample/internal/external/todolist/generated"

	"github.com/sirupsen/logrus"
)

func New(logger *logrus.Logger) external.TodoList {
	impl := &todoList{}

	wrapped := external.TodoList(impl)

	wrapped = generated.NewTodoListWithTracing(wrapped, "todolist")
	wrapped = generated.NewTodoListWithPrometheus(wrapped, "todolist")
	wrapped = generated.NewTodoListWithLogrus(wrapped, logrus.NewEntry(logger))
	wrapped = generated.NewTodoListWithTimeout(wrapped, generated.TodoListWithTimeoutConfig{
		AddTimeout:    time.Second,
		ListTimeout:   time.Second,
		RemoveTimeout: time.Second,
	})
	wrapped = generated.NewTodoListWithCircuitBreaker(wrapped, 100, time.Second, sql.ErrNoRows)
	wrapped = generated.NewTodoListWithRateLimit(wrapped, 2, 100)
	wrapped = generated.NewTodoListWithRetry(wrapped, 2, time.Second)

	return wrapped
}
