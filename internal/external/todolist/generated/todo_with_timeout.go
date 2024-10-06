// Code generated by gowrap. DO NOT EDIT.
// template: ../../../templates/timeout
// gowrap: http://github.com/hexdigest/gowrap

package generated

import (
	"context"
	_sourceExternal "goWrapExample/internal/external"
	"time"
)

// TodoListWithTimeout implements _sourceExternal.TodoList interface instrumented with timeouts
type TodoListWithTimeout struct {
	_sourceExternal.TodoList
	config TodoListWithTimeoutConfig
}

type TodoListWithTimeoutConfig struct {
	AddTimeout time.Duration

	ListTimeout time.Duration

	RemoveTimeout time.Duration
}

// NewTodoListWithTimeout returns TodoListWithTimeout
func NewTodoListWithTimeout(base _sourceExternal.TodoList, config TodoListWithTimeoutConfig) TodoListWithTimeout {
	return TodoListWithTimeout{
		TodoList: base,
		config:   config,
	}
}

// Add implements _sourceExternal.TodoList
func (_d TodoListWithTimeout) Add(ctx context.Context, text string) (s1 string, err error) {
	var cancelFunc func()
	if _d.config.AddTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.AddTimeout)
		defer cancelFunc()
	}
	return _d.TodoList.Add(ctx, text)
}

// List implements _sourceExternal.TodoList
func (_d TodoListWithTimeout) List(ctx context.Context) (sa1 []string, err error) {
	var cancelFunc func()
	if _d.config.ListTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ListTimeout)
		defer cancelFunc()
	}
	return _d.TodoList.List(ctx)
}

// Remove implements _sourceExternal.TodoList
func (_d TodoListWithTimeout) Remove(ctx context.Context, id string) (s1 string, err error) {
	var cancelFunc func()
	if _d.config.RemoveTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.RemoveTimeout)
		defer cancelFunc()
	}
	return _d.TodoList.Remove(ctx, id)
}
