package assets

import (
	"bytes"
	"context"
	_ "embed"
	"image/color"
	"main/pkg/gameerr"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed fonts/Roboto-Black.ttf
var fontBytes []byte

type Assets struct {
	Font            *text.GoTextFaceSource
	FontFace        *text.GoTextFace
	TitleFace       *text.GoTextFace
	HintFace        *text.GoTextFace
	BackgroundColor color.Color
	TextColor       color.Color
	SelectColor     color.Color
}

func LoadAssets(ctx context.Context) (*Assets, error) {
	font, err := text.NewGoTextFaceSource(bytes.NewReader(fontBytes))

	if err != nil {
		return nil, gameerr.New(
			gameerr.ErrCodeAssetsLoading,
			"failed to load font face from assets",
			err,
			gameerr.SeverityFatal,
			map[string]any{"attempt": 1},
		)
	}

	return &Assets{
		Font:            font,
		FontFace:        &text.GoTextFace{Source: font, Size: 32},
		TitleFace:       &text.GoTextFace{Source: font, Size: 48},
		HintFace:        &text.GoTextFace{Source: font, Size: 16},
		BackgroundColor: color.RGBA{R: 20, G: 20, B: 40, A: 255},
		TextColor:       color.RGBA{R: 180, G: 180, B: 180, A: 255},
		SelectColor:     color.RGBA{R: 255, G: 220, B: 100, A: 255},
	}, nil
}
