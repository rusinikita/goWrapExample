package todolist

import (
	"context"
)

type todoList struct {
}

func (t *todoList) List(ctx context.Context) ([]string, error) {
	return nil, nil
}

func (t *todoList) Add(ctx context.Context, text string) (string, error) {
	return "", nil
}

func (t *todoList) Remove(ctx context.Context, id string) (string, error) {
	return "", nil
}
