package game

import (
	"context"

	"github.com/hajimehoshi/ebiten/v2"
)

type PlayState struct {
}

func NewPlayState() *PlayState {
	return &PlayState{}
}

func (p *PlayState) Update(g *Game) error {
	//TODO implement me
	panic("implement me")
}

func (p *PlayState) Draw(g *Game, screen *ebiten.Image) {
	//TODO implement me
	panic("implement me")
}

func (p *PlayState) OnEnter(ctx context.Context, g *Game) {
	//TODO implement me
	panic("implement me")
}

func (p *PlayState) OnExit(ctx context.Context, g *Game) {
	//TODO implement me
	panic("implement me")
}
