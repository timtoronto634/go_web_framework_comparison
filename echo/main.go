package main

import "golang.org/x/exp/slog"

func main() {
	s := server()
	err := s.Start(":1323")
	if err != nil {
		slog.Error(err.Error())
	}
}
