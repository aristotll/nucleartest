package main

import (
	"errors"
	"golang.org/x/exp/slog"
	"io"
	"os"
)

type S struct {
	a, b, c int64
	x, y, z string
}

func main() {
	s := &S{1, 2, 3, "4", "5", "6"}
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout)))
	slog.Info("msg", slog.String("name", "zhang3"))
	slog.Info("hello", "name", "Al")

	errval := errors.New("not found this user")
	slog.Error("not found this user", errval, slog.String("userID", "1001"))
	slog.Error("msg", io.EOF, slog.Any("struct", s))
}
