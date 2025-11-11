package main

import (
	"log/slog"
	"os"
	"testing"
)

func TestSLog(t *testing.T) {
	h := slog.NewTextHandler(os.Stderr, nil)
	slog.SetDefault(slog.New(h))
	slog.Info("hello", "count", 3)
}
