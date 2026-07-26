package game

import (
	"context"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type ItemMenu struct {
	Label   string
	OnClick func(g *Game)
}

type MenuState struct {
	items         []ItemMenu
	selectedIdx   int
	inputCooldown int
}

func NewMenuState() *MenuState {
	m := &MenuState{}

	m.items = []ItemMenu{
		{
			Label: "Start Game",
			OnClick: func(g *Game) {
				g.ChangeState("play")
			},
		},
		{
			Label: "Settings",
			OnClick: func(g *Game) {
				g.ChangeState("settings")
			},
		},
		{
			Label: "Quit",
			OnClick: func(g *Game) {
				os.Exit(0)
			},
		},
	}

	return m
}

func (m *MenuState) OnEnter(ctx context.Context, g *Game) {
	m.selectedIdx = 0
	m.inputCooldown = InputCooldown
}

func (m *MenuState) OnExit(ctx context.Context, g *Game) {}

func (m *MenuState) Update(g *Game) error {
	// Кулдаун, чтобы нажатие клавиши не обрабатывалось каждый кадр
	if m.inputCooldown > 0 {
		m.inputCooldown--
	} else {
		// Навигация вверх/вниз
		if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
			m.selectedIdx--
			if m.selectedIdx < 0 {
				m.selectedIdx = len(m.items) - 1
			}
			m.inputCooldown = InputCooldown
		}

		if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
			m.selectedIdx++
			if m.selectedIdx >= len(m.items) {
				m.selectedIdx = 0
			}
			m.inputCooldown = InputCooldown
		}

		// Выбор пункта
		if ebiten.IsKeyPressed(ebiten.KeyEnter) || ebiten.IsKeyPressed(ebiten.KeySpace) {
			m.items[m.selectedIdx].OnClick(g)
			m.inputCooldown = InputCooldown
		}
	}

	return nil
}

func (m *MenuState) Draw(g *Game, screen *ebiten.Image) {
	// Фон
	screen.Fill(color.RGBA{R: 20, G: 20, B: 40, A: 255})

	// Заголовок
	titleOp := &text.DrawOptions{}
	titleOp.GeoM.Translate(float64(g.ScreenWidth/2), 100)
	titleOp.ColorScale.Scale(1, 1, 1, 1) // белый
	titleOp.PrimaryAlign = text.AlignCenter
	titleOp.SecondaryAlign = text.AlignCenter

	titleFace := &text.GoTextFace{
		Source: g.Assets.Font,
		Size:   48,
	}
	text.Draw(screen, "WSG", titleFace, titleOp)

	// Пункты меню
	startY := 250
	itemSpacing := 80

	for i, item := range m.items {
		y := startY + i*itemSpacing

		// Цвет текста
		textColor := color.RGBA{R: 180, G: 180, B: 180, A: 255} // серый
		prefix := "  "

		if i == m.selectedIdx {
			textColor = color.RGBA{R: 255, G: 220, B: 100, A: 255} // золотистый
			prefix = "> "
		}

		// Рисуем пункт
		op := &text.DrawOptions{}
		op.GeoM.Translate(400, float64(y))
		op.ColorScale.ScaleWithColor(textColor)
		op.PrimaryAlign = text.AlignCenter
		op.SecondaryAlign = text.AlignCenter

		text.Draw(screen, prefix+item.Label, g.Assets.FontFace, op)
	}

	// Подсказка внизу
	hintOp := &text.DrawOptions{}
	hintOp.GeoM.Translate(400, 550)
	hintOp.ColorScale.Scale(0.5, 0.5, 0.5, 1)
	hintOp.PrimaryAlign = text.AlignCenter
	hintOp.SecondaryAlign = text.AlignCenter

	smallFace := &text.GoTextFace{
		Source: g.Assets.Font,
		Size:   16,
	}
	text.Draw(screen, "Use up/down arrow keys   |   ENTER - confirm", smallFace, hintOp)
}
