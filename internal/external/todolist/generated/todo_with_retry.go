// Code generated by gowrap. DO NOT EDIT.
// template: ../../../templates/retry
// gowrap: http://github.com/hexdigest/gowrap

package generated

import (
	"context"
	_sourceExternal "goWrapExample/internal/external"
	"time"
)

// TodoListWithRetry implements _sourceExternal.TodoList interface instrumented with retries
type TodoListWithRetry struct {
	_sourceExternal.TodoList
	_retryCount    int
	_retryInterval time.Duration
}

// NewTodoListWithRetry returns TodoListWithRetry
func NewTodoListWithRetry(base _sourceExternal.TodoList, retryCount int, retryInterval time.Duration) TodoListWithRetry {
	return TodoListWithRetry{
		TodoList:       base,
		_retryCount:    retryCount,
		_retryInterval: retryInterval,
	}
}

// Add implements _sourceExternal.TodoList
func (_d TodoListWithRetry) Add(ctx context.Context, text string) (s1 string, err error) {
	s1, err = _d.TodoList.Add(ctx, text)
	if err == nil || _d._retryCount < 1 {
		return
	}
	_ticker := time.NewTicker(_d._retryInterval)
	defer _ticker.Stop()
	for _i := 0; _i < _d._retryCount && err != nil; _i++ {
		select {
		case <-ctx.Done():
			return
		case <-_ticker.C:
		}
		s1, err = _d.TodoList.Add(ctx, text)
	}
	return
}

// List implements _sourceExternal.TodoList
func (_d TodoListWithRetry) List(ctx context.Context) (sa1 []string, err error) {
	sa1, err = _d.TodoList.List(ctx)
	if err == nil || _d._retryCount < 1 {
		return
	}
	_ticker := time.NewTicker(_d._retryInterval)
	defer _ticker.Stop()
	for _i := 0; _i < _d._retryCount && err != nil; _i++ {
		select {
		case <-ctx.Done():
			return
		case <-_ticker.C:
		}
		sa1, err = _d.TodoList.List(ctx)
	}
	return
}

// Remove implements _sourceExternal.TodoList
func (_d TodoListWithRetry) Remove(ctx context.Context, id string) (s1 string, err error) {
	s1, err = _d.TodoList.Remove(ctx, id)
	if err == nil || _d._retryCount < 1 {
		return
	}
	_ticker := time.NewTicker(_d._retryInterval)
	defer _ticker.Stop()
	for _i := 0; _i < _d._retryCount && err != nil; _i++ {
		select {
		case <-ctx.Done():
			return
		case <-_ticker.C:
		}
		s1, err = _d.TodoList.Remove(ctx, id)
	}
	return
}
