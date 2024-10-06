package main

import (
	"context"

	"goWrapExample/internal/external/todolist"

	"github.com/sirupsen/logrus"
)

func main() {
	client := todolist.New(logrus.New())

	_, _ = client.List(context.Background())
}
