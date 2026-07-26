package game

import (
	"context"

	"github.com/hajimehoshi/ebiten/v2"
)

type State interface {
	Update(g *Game) error
	Draw(g *Game, screen *ebiten.Image)
	OnEnter(ctx context.Context, g *Game)
	OnExit(ctx context.Context, g *Game)
}

type StateManager struct {
	states map[string]State
}

func NewStateManager() *StateManager {
	return &StateManager{
		states: map[string]State{
			"menu":     NewMenuState(),
			"play":     NewPlayState(),
			"settings": NewSettingsState(),
		},
	}
}

func (s *StateManager) Get(name string) State {
	return s.states[name]
}
