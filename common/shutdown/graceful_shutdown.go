package shutdown

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func GracefulShutdown() context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-exitSignal
		cancel()
	}()

	return ctx
}
