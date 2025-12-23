package application

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func Run(ctx context.Context) {
	var waitGroup sync.WaitGroup

	waitGroup.Add(1)

	process := make(chan os.Signal, 1)

	signal.Notify(process, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer waitGroup.Done()

		select {
		case <-ctx.Done():
			slog.Info("Application context cancelled")
			return
		case <-process:
			slog.Info("Application interrupted")
			return
		}
	}()

	waitGroup.Wait()

	slog.Info("Application stopped")
}
