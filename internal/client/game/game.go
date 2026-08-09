package game

import (
	"context"
	"errors"
	"fmt"
	"main/internal/client/assets"
	"main/internal/client/game/settings"
	"main/pkg/gameerr"
	"main/pkg/logger"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	assets       *assets.Assets
	currentState State
	stateManager *StateManager
	settings     *settings.Settings
	screenWidth  int
	screenHeight int
	ctx          context.Context
	cancel       context.CancelFunc
}

func NewGame(ctx context.Context, width, height int) (*Game, error) {
	ctx, cancel := context.WithCancel(ctx)

	a, err := assets.LoadAssets(ctx)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("load assets err: %w", err)
	}

	newSettings, err := settings.NewSettings(ctx)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("load settings err: %w", err)
	}

	g := &Game{
		assets:       a,
		stateManager: NewStateManager(),
		settings:     newSettings,
		screenWidth:  width,
		screenHeight: height,
		ctx:          ctx,
		cancel:       cancel,
	}

	logger.Info(g.ctx, "Client side game initialized")

	if err := g.ChangeState(g.ctx, StateMenuKey); err != nil {
		return nil, fmt.Errorf("set menu state fail: %w", err)
	}

	return g, nil
}

func (g *Game) ChangeState(ctx context.Context, state StateKey) error {
	next, ok := g.stateManager.Get(state)
	if !ok {
		return gameerr.New(
			gameerr.ErrCodeStateUnknown,
			"failed to change state",
			nil,
			gameerr.SeverityFatal,
			map[string]any{
				"state": state,
			},
		)
	}

	logger.Info(ctx, "Changing game state", "new_state", state)

	if g.currentState != nil {
		g.currentState.OnExit(ctx, g)
	}

	g.currentState = next
	g.currentState.OnEnter(ctx, g)

	logger.SetClientState(string(state))

	return nil
}

// Context returns the current game context.
func (g *Game) Context() context.Context {
	return g.ctx
}

// WithContext allows you to dynamically update the context.
func (g *Game) WithContext(ctx context.Context) {
	if ctx == nil {
		panic("nil context")
	}
	g.ctx = ctx
}

func (g *Game) HandleError(ctx context.Context, err error) {
	if err == nil {
		return
	}

	var gErr *gameerr.GameError

	if errors.As(err, &gErr) {
		switch gErr.Severity {
		case gameerr.SeverityFatal:
			gErr.Log(ctx)
			g.Cancel()
		default:
			gErr.Log(ctx)
		}

		return
	}

	logger.Warn(ctx, "Unknown error occurred", "error", err)
}

// Cancel soft stop of the game.
func (g *Game) Cancel() {
	g.cancel()
}

func (g *Game) Update() error {
	select {
	case <-g.ctx.Done():
		return ebiten.Termination
	default:
		g.currentState.Update(g)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.currentState.Draw(g, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}
