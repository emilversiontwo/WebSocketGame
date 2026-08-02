package game

import (
	"context"
	"fmt"
	"main/internal/client/game/settings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type SettingsItem struct {
	Label       string
	IsInput     bool
	SettingName string
	OnClick     func(item SettingsItem, s *SettingsState, g *Game)
}

type SettingsState struct {
	items         []SettingsItem
	settingsCache map[string]any
	selectedIdx   int
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
			SettingName: string(settings.KeyPlayerName),
			OnClick: func(item SettingsItem, s *SettingsState, g *Game) {
				s.typeMode = true
				s.text = s.settingsCache[item.SettingName].(string)
			},
		},
		{
			Label:       "Change Color",
			IsInput:     true,
			SettingName: string(settings.KeyPlayerColor),
			OnClick: func(item SettingsItem, s *SettingsState, g *Game) {
				s.typeMode = true
				s.text = s.settingsCache[item.SettingName].(string)
			},
		},
		{
			Label:       "Change Language",
			IsInput:     true,
			SettingName: string(settings.KeyLocale),
			OnClick: func(item SettingsItem, s *SettingsState, g *Game) {
				if s.settingsCache[item.SettingName].(string) == "ru_RU" {
					s.settingsCache[item.SettingName] = "en_US"
					return
				}
				s.settingsCache[item.SettingName] = "ru_RU"
			},
		},
		{
			Label:   "Save and Back",
			IsInput: false,
			OnClick: func(item SettingsItem, s *SettingsState, g *Game) {
				if err := g.ChangeState(g.ctx, StateMenuKey); err != nil {
					g.Cancel()
				}
			},
		},
		{
			Label:   "Restore",
			IsInput: false,
			OnClick: func(item SettingsItem, s *SettingsState, g *Game) {
				settingsMap, err := g.settings.ToMap(g.ctx)
				if err != nil {
					g.Cancel()
				}
				s.settingsCache = settingsMap
			},
		},
	}

	return s
}

func (s *SettingsState) Update(g *Game) {
	if s.typeMode {
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			s.typeMode = false
			s.settingsCache[s.items[s.selectedIdx].SettingName] = s.text
		}

		s.runes = ebiten.AppendInputChars(s.runes[:0])
		s.text += string(s.runes)

		if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
			if len(s.text) >= 1 {
				s.text = s.text[:len(s.text)-1]
			}
		}
	} else {
		if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
			s.selectedIdx--
			if s.selectedIdx < 0 {
				s.selectedIdx = len(s.items) - 1
			}
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
			s.selectedIdx++
			if s.selectedIdx >= len(s.items) {
				s.selectedIdx = 0
			}
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			item := s.items[s.selectedIdx]
			item.OnClick(item, s, g)
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			if err := g.ChangeState(g.ctx, StateMenuKey); err != nil {
				g.Cancel()
			}
		}
	}
}

func (s *SettingsState) Draw(g *Game, screen *ebiten.Image) {
	screen.Fill(g.assets.BackgroundColor)

	textOp := &text.DrawOptions{}
	textOp.GeoM.Translate(float64(g.screenWidth/2), 100)
	textOp.ColorScale.ScaleWithColor(g.assets.TextColor)
	textOp.PrimaryAlign = text.AlignCenter
	textOp.SecondaryAlign = text.AlignCenter

	text.Draw(screen, "WSG", g.assets.TitleFace, textOp)

	textOp.GeoM.Translate(0, 50)

	text.Draw(screen, "Settings", g.assets.FontFace, textOp)

	startY := 250
	itemSpacing := 80

	for i, item := range s.items {
		y := startY + i*itemSpacing

		textColor := g.assets.TextColor
		prefix := "  "

		if i == s.selectedIdx {
			textColor = g.assets.SelectColor
			prefix = "> "
		}
		if i == s.selectedIdx && item.IsInput && s.typeMode {
			vector.FillRect(
				screen,
				float32(g.screenWidth/4), float32(y)-40,
				float32(20*len(s.text)), 80,
				g.assets.TextColor,
				true,
			)

			op := &text.DrawOptions{}
			op.GeoM.Translate(float64(g.screenWidth/3), float64(y))
			op.ColorScale.ScaleWithColor(textColor)
			op.PrimaryAlign = text.AlignStart
			op.SecondaryAlign = text.AlignCenter
			t := s.text
			if ebiten.Tick()%60 < 30 {
				t += "_"
			}

			text.Draw(screen, prefix+t, g.assets.FontFace, op)
		} else {
			op := &text.DrawOptions{}
			op.GeoM.Translate(float64(g.screenWidth/3), float64(y))
			op.ColorScale.ScaleWithColor(textColor)
			op.PrimaryAlign = text.AlignCenter
			op.SecondaryAlign = text.AlignCenter

			text.Draw(screen, prefix+item.Label, g.assets.FontFace, op)
			if item.IsInput {
				op.GeoM.Translate(float64(g.screenWidth/5), 0)
				op.PrimaryAlign = text.AlignStart

				text.Draw(screen, fmt.Sprintf(" -> %v", s.settingsCache[item.SettingName]), g.assets.FontFace, op)
			}
		}
	}
}

func (s *SettingsState) OnEnter(ctx context.Context, g *Game) {
	s.selectedIdx = 0
	s.typeMode = false
	s.text = ""
	settingsMap, err := g.settings.ToMap(ctx)
	if err != nil {
		g.Cancel()
	}
	s.settingsCache = settingsMap
}

func (s *SettingsState) OnExit(ctx context.Context, g *Game) {

	if g.settings.SetFromMap(ctx, s.settingsCache) != nil {
		g.Cancel()
	}

	if g.settings.Save(ctx) != nil {
		g.Cancel()
	}
}
