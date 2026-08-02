package game

import (
	"context"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type ItemMenu struct {
	Label   string
	OnClick func(g *Game)
}

type MenuState struct {
	items       []ItemMenu
	selectedIdx int
}

func NewMenuState() *MenuState {
	m := &MenuState{}

	m.items = []ItemMenu{
		{
			Label: "Start Game",
			OnClick: func(g *Game) {
				if err := g.ChangeState(g.ctx, StatePlayKey); err != nil {
					g.Cancel()
				}
			},
		},
		{
			Label: "Settings",
			OnClick: func(g *Game) {
				if err := g.ChangeState(g.ctx, StateSettingsKey); err != nil {
					g.Cancel()
				}
			},
		},
		{
			Label: "Quit",
			OnClick: func(g *Game) {
				g.Cancel()
			},
		},
	}

	return m
}

func (m *MenuState) OnEnter(ctx context.Context, g *Game) {
	m.selectedIdx = 0
}

func (m *MenuState) OnExit(ctx context.Context, g *Game) {}

func (m *MenuState) Update(g *Game) {
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		m.selectedIdx--
		if m.selectedIdx < 0 {
			m.selectedIdx = len(m.items) - 1
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		m.selectedIdx++
		if m.selectedIdx >= len(m.items) {
			m.selectedIdx = 0
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		m.items[m.selectedIdx].OnClick(g)
	}
}

func (m *MenuState) Draw(g *Game, screen *ebiten.Image) {
	screen.Fill(g.assets.BackgroundColor)

	titleOp := &text.DrawOptions{}
	titleOp.GeoM.Translate(float64(g.screenWidth/2), 100)
	titleOp.ColorScale.ScaleWithColor(g.assets.TextColor)
	titleOp.PrimaryAlign = text.AlignCenter
	titleOp.SecondaryAlign = text.AlignCenter

	text.Draw(screen, "WSG", g.assets.TitleFace, titleOp)

	startY := 250
	itemSpacing := 80

	for i, item := range m.items {
		y := startY + i*itemSpacing

		textColor := g.assets.TextColor
		prefix := "  "

		if i == m.selectedIdx {
			textColor = g.assets.SelectColor
			prefix = "> "
		}

		op := &text.DrawOptions{}
		op.GeoM.Translate(400, float64(y))
		op.ColorScale.ScaleWithColor(textColor)
		op.PrimaryAlign = text.AlignCenter
		op.SecondaryAlign = text.AlignCenter

		text.Draw(screen, prefix+item.Label, g.assets.FontFace, op)
	}

	hintOp := &text.DrawOptions{}
	hintOp.GeoM.Translate(400, 550)
	hintOp.ColorScale.ScaleWithColor(g.assets.TextColor)
	hintOp.PrimaryAlign = text.AlignCenter
	hintOp.SecondaryAlign = text.AlignCenter

	text.Draw(screen, "Use up/down arrow keys   |   ENTER - confirm", g.assets.HintFace, hintOp)
}
