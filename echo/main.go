package main

import (
	"fmt"
	"timtoronto634/go_web_framework_comparison/echo/router"

	"github.com/timtoronto634/go_web_framework_comparison/db"
	"golang.org/x/exp/slog"
)

func main() {
	db, closer, err := db.NewDB()
	if err != nil {
		fmt.Println("echo: no db", err)
		return
	}
	defer closer()

	if err := db.Ping(); err != nil {
		fmt.Println("echo: no resposne for ping", err)
		return
	}

	s := router.Server()
	err = s.Start(":1323")
	if err != nil {
		slog.Error(err.Error())
	}
}
