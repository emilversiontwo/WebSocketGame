package game

import (
	"context"
	"log/slog"
	"main/internal/client/assets"
	"main/internal/client/game/settings"
	"main/pkg/logger"

	"github.com/hajimehoshi/ebiten/v2"
)

const InputCooldown int = 10

type Game struct {
	Assets       *assets.Assets
	CurrentState State
	StateManager *StateManager
	Settings     *settings.Settings
	ScreenWidth  int
	ScreenHeight int
	Ctx          context.Context
}

func NewGame(ctx context.Context, width, height int) *Game {
	g := &Game{
		Assets:       assets.LoadAssets(ctx),
		StateManager: NewStateManager(),
		Settings:     settings.NewSettings(ctx),
		ScreenWidth:  width,
		ScreenHeight: height,
		Ctx:          ctx,
	}

	slog.InfoContext(g.Ctx, "Client side game initialized")

	g.ChangeState("menu")

	return g
}

func (g *Game) ChangeState(state string) {
	slog.InfoContext(g.Ctx, "Changing game state", "new_state", state)
	if g.CurrentState != nil {
		g.CurrentState.OnExit(g.Ctx, g)
	}
	g.CurrentState = g.StateManager.Get(state)
	g.CurrentState.OnEnter(g.Ctx, g)
	logger.SetClientState(state)
	slog.InfoContext(g.Ctx, "Changed game state")
}

func (g *Game) Update() error {
	return g.CurrentState.Update(g)
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.CurrentState.Draw(g, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}
