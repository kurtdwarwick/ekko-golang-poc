package application

import (
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	process := make(chan os.Signal, 1)

	signal.Notify(process, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-process
}
