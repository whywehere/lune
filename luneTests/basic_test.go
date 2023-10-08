package main

import (
	"github.com/whywehere/lune"
	"log/slog"
	"net/http"
	"testing"
)

func test(http.ResponseWriter, *http.Request) {
	slog.Info("Testing")
}

func TestBasic(t *testing.T) {
	e := lune.New()
	e.POST("/test", test)
	err := e.Run(":8080")
	if err != nil {
		slog.Error("Failed to start lune")
		return
	}
}
