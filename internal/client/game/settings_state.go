package game

import (
	"context"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type SettingsItem struct {
	Label       string
	IsInput     bool
	SettingName string
	OnClick     func(s *SettingsState, g *Game)
}

type SettingsState struct {
	items         []SettingsItem
	settingsCache map[string]any
	selectedIdx   int
	inputCooldown int
	typeMode      bool
	text          string
	runes         []rune
}

func NewSettingsState() *SettingsState {
	s := &SettingsState{}

	s.items = []SettingsItem{
		{
			Label:       "Change Name",
			IsInput:     true,
			SettingName: "PlayerName",
			OnClick: func(s *SettingsState, g *Game) {
				s.typeMode = true
				s.text = s.settingsCache["PlayerName"].(string)
			},
		},
		{
			Label:       "Change Color",
			IsInput:     true,
			SettingName: "PlayerColor",
			OnClick: func(s *SettingsState, g *Game) {
				s.typeMode = true
				s.text = s.settingsCache["PlayerColor"].(string)
			},
		},
		{
			Label:       "Change Language",
			IsInput:     true,
			SettingName: "Locale",
			OnClick: func(s *SettingsState, g *Game) {
				if s.settingsCache["Locale"].(string) == "ru_RU" {
					s.settingsCache["Locale"] = "en_US"
					return
				}
				s.settingsCache["Locale"] = "ru_RU"
			},
		},
		{
			Label:   "Save and Back",
			IsInput: false,
			OnClick: func(s *SettingsState, g *Game) {
				g.ChangeState("menu")
			},
		},
		{
			Label:   "Restore",
			IsInput: false,
			OnClick: func(s *SettingsState, g *Game) {
				s.settingsCache = g.Settings.GetSettingsMap(nil)
			},
		},
	}

	return s
}

func (s *SettingsState) Update(g *Game) error {
	if s.inputCooldown > 0 {
		s.inputCooldown--
	} else {
		if s.typeMode {
			if ebiten.IsKeyPressed(ebiten.KeyEnter) {
				s.typeMode = false
				s.inputCooldown = InputCooldown + 5
				s.settingsCache[s.items[s.selectedIdx].SettingName] = s.text
			}

			s.runes = ebiten.AppendInputChars(s.runes[:0])
			s.text += string(s.runes)

			if ebiten.IsKeyPressed(ebiten.KeyBackspace) {
				if len(s.text) >= 1 {
					s.text = s.text[:len(s.text)-1]
				}
				s.inputCooldown = 3
			}
		} else {
			if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
				s.selectedIdx--
				if s.selectedIdx < 0 {
					s.selectedIdx = len(s.items) - 1
				}
				s.inputCooldown = InputCooldown
			}

			if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
				s.selectedIdx++
				if s.selectedIdx >= len(s.items) {
					s.selectedIdx = 0
				}
				s.inputCooldown = InputCooldown
			}

			if ebiten.IsKeyPressed(ebiten.KeyEnter) || ebiten.IsKeyPressed(ebiten.KeySpace) {
				s.items[s.selectedIdx].OnClick(s, g)
				s.inputCooldown = InputCooldown + 5
			}

			if ebiten.IsKeyPressed(ebiten.KeyEscape) {
				g.ChangeState("menu")
			}
		}
	}

	return nil
}

func (s *SettingsState) Draw(g *Game, screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 20, G: 20, B: 40, A: 255})

	textOp := &text.DrawOptions{}
	textOp.GeoM.Translate(float64(g.ScreenWidth/2), 100)
	textOp.ColorScale.Scale(1, 1, 1, 1) // белый
	textOp.PrimaryAlign = text.AlignCenter
	textOp.SecondaryAlign = text.AlignCenter

	titleFace := &text.GoTextFace{
		Source: g.Assets.Font,
		Size:   48,
	}
	text.Draw(screen, "WSG", titleFace, textOp)

	textOp.GeoM.Translate(0, 50)

	textUnderTitleFace := &text.GoTextFace{
		Source: g.Assets.Font,
		Size:   24,
	}
	text.Draw(screen, "Settings", textUnderTitleFace, textOp)

	startY := 250
	itemSpacing := 80

	for i, item := range s.items {
		y := startY + i*itemSpacing

		textColor := color.RGBA{R: 180, G: 180, B: 180, A: 255} // серый
		prefix := "  "

		if i == s.selectedIdx {
			textColor = color.RGBA{R: 255, G: 220, B: 100, A: 255} // золотистый
			prefix = "> "
		}
		if i == s.selectedIdx && item.IsInput && s.typeMode {
			vector.FillRect(
				screen,
				float32(g.ScreenWidth/4), float32(y)-40,
				float32(20*len(s.text)), 80,
				color.RGBA{255, 255, 255, 255},
				true,
			)

			op := &text.DrawOptions{}
			op.GeoM.Translate(float64(g.ScreenWidth/3), float64(y))
			op.ColorScale.ScaleWithColor(textColor)
			op.PrimaryAlign = text.AlignStart
			op.SecondaryAlign = text.AlignCenter
			t := s.text
			if ebiten.Tick()%60 < 30 {
				t += "_"
			}

			text.Draw(screen, prefix+t, g.Assets.FontFace, op)
		} else {
			op := &text.DrawOptions{}
			op.GeoM.Translate(float64(g.ScreenWidth/3), float64(y))
			op.ColorScale.ScaleWithColor(textColor)
			op.PrimaryAlign = text.AlignCenter
			op.SecondaryAlign = text.AlignCenter

			text.Draw(screen, prefix+item.Label, g.Assets.FontFace, op)
			if item.IsInput {
				op.GeoM.Translate(float64(g.ScreenWidth/5), 0)
				op.PrimaryAlign = text.AlignStart
				textValueFace := &text.GoTextFace{
					Source: g.Assets.Font,
					Size:   24,
				}

				text.Draw(screen, fmt.Sprintf(" -> %v", s.settingsCache[item.SettingName]), textValueFace, op)
			}
		}
	}
}

func (s *SettingsState) OnEnter(ctx context.Context, g *Game) {
	s.selectedIdx = 0
	s.inputCooldown = InputCooldown
	s.typeMode = false
	s.text = ""
	s.settingsCache = g.Settings.GetSettingsMap(nil)
}

func (s *SettingsState) OnExit(ctx context.Context, g *Game) {
	g.Settings.SetSettingsFromMap(ctx, s.settingsCache)
}
