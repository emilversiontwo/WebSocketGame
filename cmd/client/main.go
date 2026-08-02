package main

import (
	"context"
	"log/slog"
	"main/internal/client/game"
	"main/pkg/logger"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	logger.InitClient(slog.LevelDebug, "0.0.1", "WSG_client")

	ctx := context.Background()

	slog.InfoContext(ctx, "Starting Client")

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("WSG")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	g, err := game.NewGame(ctx, 800, 600)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to initialize Game")
		panic(err)
	}
	defer g.Cancel()

	if err := ebiten.RunGame(g); err != nil {
		slog.ErrorContext(ctx, "Client stopped with error", "err", err)
	}
}
