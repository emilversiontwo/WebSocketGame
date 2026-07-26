package assets

import (
	"bytes"
	"context"
	_ "embed"
	"log/slog"
	"main/pkg/gameerr"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed fonts/Roboto-Black.ttf
var fontBytes []byte

type Assets struct {
	Font     *text.GoTextFaceSource
	FontFace *text.GoTextFace
}

func LoadAssets(ctx context.Context) *Assets {
	font, err := text.NewGoTextFaceSource(bytes.NewReader(fontBytes))

	if err != nil {
		err := gameerr.New(
			"ERR_FILE_LOADING",
			"failed to load font face from assets",
			err,
			map[string]any{
				"attempt": 1,
			},
		)
		err.Log(ctx, slog.LevelError)
	}

	return &Assets{
		Font: font,
		FontFace: &text.GoTextFace{
			Source: font,
			Size:   32,
		},
	}
}
