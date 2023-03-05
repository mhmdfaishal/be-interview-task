package main

import (
	"be-interview-task/internal/boot"
	"be-interview-task/internal/server"
)

func main() {
	defer boot.FlushResources()

	s := server.NewServer()
	s.Start()
}
