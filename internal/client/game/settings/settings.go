package settings

import (
	"context"
	"encoding/json"
	"log/slog"
	"main/pkg/ctxutil"
	"main/pkg/gameerr"
	"math/rand"
	"os"
	"strconv"
)

const settingsFileName string = "settings.json"

type Settings struct {
	PlayerName string `json:"player_name"`
	// PlayerColor is HEX
	PlayerColor string `json:"player_color"`
	Locale      string `json:"locale"`
}

func (s Settings) saveSettingsAsync(ctx context.Context) {
	go func() {
		data, err := json.MarshalIndent(s, "", "  ")
		if err != nil {
			gameerr.New("ERR_SETTINGS_MARSHAL", "failed to marshal settings", err, nil).LogAndReturn(ctx, slog.LevelError)
			return
		}

		err = os.WriteFile(settingsFileName, data, 0644)
		if err != nil {
			gameerr.New("ERR_SETTINGS_WRITE", "failed to write settings file", err, map[string]any{
				"filename": settingsFileName,
			}).LogAndReturn(ctx, slog.LevelError)
			return
		}

		slog.InfoContext(ctx, "Settings saved successfully", "filename", settingsFileName)
	}()
}

func (s *Settings) SetSettingsFromMap(ctx context.Context, settings map[string]any) {
	//TODO: Придумать механизм мапинга значений в обе стороны
	for k, v := range settings {
		switch k {
		case "PlayerName":
			s.PlayerName = v.(string)
		case "PlayerColor":
			s.PlayerColor = v.(string)
		case "Locale":
			s.Locale = v.(string)
		default:
			err := gameerr.New(
				"ERR_SETTINGS_MAPING",
				"failed to map settings",
				nil,
				map[string]any{
					"key":   k,
					"value": v,
				},
			)
			err.Log(ctx, slog.LevelError)
		}
	}

	s.saveSettingsAsync(ctx)
}

func (s *Settings) GetSettingsMap(name *string) map[string]any {
	settingsMap := map[string]any{
		"PlayerName":  s.PlayerName,
		"PlayerColor": s.PlayerColor,
		"Locale":      s.Locale,
	}

	if name == nil {
		return settingsMap
	}

	for key, value := range settingsMap {
		if key == *name {
			return map[string]any{
				key: value,
			}
		}
	}

	return settingsMap
}

func NewSettings(ctx context.Context) *Settings {
	file, err := os.Open(settingsFileName)

	if err != nil {
		err := gameerr.New(
			"ERR_FILE_LOADING",
			"failed to load settings file is missing",
			err,
			map[string]any{
				"filename": settingsFileName,
				"attempt":  1,
			},
		)
		err.Log(ctx, slog.LevelError)

		return createDefaultSettings(ctx)
	}

	defer func(file *os.File) {
		err := file.Close()

		if err != nil {
			err := gameerr.New(
				"ERR_FILE_CLOSING",
				"failed to close settings file",
				err,
				map[string]any{
					"filename": settingsFileName,
					"attempt":  1,
				},
			)
			err.Log(ctx, slog.LevelError)
		}
	}(file)

	decoder := json.NewDecoder(file)

	var settings Settings

	if err := decoder.Decode(&settings); err != nil {
		err := gameerr.New(
			"ERR_FILE_READING",
			"failed to reading settings file, it is bad",
			err,
			map[string]any{
				"filename": settingsFileName,
				"attempt":  1,
			},
		)
		err.Log(ctx, slog.LevelError)

		return createDefaultSettings(ctx)
	}

	ctxutil.WithLocale(ctx, settings.Locale)
	slog.InfoContext(ctx, "Settings file is loaded")
	return &settings
}

func createDefaultSettings(ctx context.Context) *Settings {
	settings := &Settings{
		PlayerName:  "Player" + strconv.Itoa(rand.Intn(100)),
		PlayerColor: "DC143C",
		Locale:      "en_US",
	}

	ctxutil.WithLocale(ctx, settings.Locale)

	return settings
}
