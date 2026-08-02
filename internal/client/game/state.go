package game

import (
	"context"

	"github.com/hajimehoshi/ebiten/v2"
)

type StateKey string

const (
	StateMenuKey     StateKey = "menu"
	StatePlayKey     StateKey = "play"
	StateSettingsKey StateKey = "settings"
)

var registry = map[StateKey]State{
	StateMenuKey:     NewMenuState(),
	StatePlayKey:     NewPlayState(),
	StateSettingsKey: NewSettingsState(),
}

type State interface {
	Update(g *Game)
	Draw(g *Game, screen *ebiten.Image)
	OnEnter(ctx context.Context, g *Game)
	OnExit(ctx context.Context, g *Game)
}

type StateManager struct {
	states map[StateKey]State
}

func NewStateManager() *StateManager {
	return &StateManager{
		states: registry,
	}
}

func (s *StateManager) Get(name StateKey) (State, bool) {
	st, ok := s.states[name]
	return st, ok
}
