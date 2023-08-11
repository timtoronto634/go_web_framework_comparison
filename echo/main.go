package main

import (
	"timtoronto634/go_web_framework_comparison/echo/router"

	"golang.org/x/exp/slog"
)

func main() {
	s := router.Server()
	err := s.Start(":1323")
	if err != nil {
		slog.Error(err.Error())
	}
}
