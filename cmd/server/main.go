package main

import (
	"context"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	container := NewContainer()
	if err := container.Start(ctx); err != nil {
		container.logger.Fatal().Err(err).Msg("failed to start the server")
	}
}
