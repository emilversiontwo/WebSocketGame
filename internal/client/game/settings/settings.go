package settings

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"main/pkg/gameerr"
	"math/rand"
	"os"
	"sync"
)

type SettingKey string

const (
	KeyPlayerName  SettingKey = "player_name"
	KeyPlayerColor SettingKey = "player_color"
	KeyLocale      SettingKey = "locale"
)

const settingsFileName string = "settings.json"

type Settings struct {
	mu sync.RWMutex

	PlayerName string `json:"player_name"`
	// PlayerColor is HEX
	PlayerColor string `json:"player_color"`
	Locale      string `json:"locale"`
}

// Save the settings to a file.
func (s *Settings) Save(ctx context.Context) gameerr.GameErrorer {
	s.mu.RLock()
	defer s.mu.RUnlock()

	data, err := json.MarshalIndent(s, "", "  ")

	if err != nil {
		return gameerr.New(
			gameerr.ErrCodeSettingsMarshal,
			"failed to marshal settings",
			err,
			nil,
		).LogAndReturn(ctx, slog.LevelError)
	}

	if err = os.WriteFile(settingsFileName, data, 0644); err != nil {
		return gameerr.New(
			gameerr.ErrCodeSettingsWriting,
			"failed to write settings file",
			err,
			map[string]any{"filename": settingsFileName},
		).LogAndReturn(ctx, slog.LevelError)
	}

	return nil
}

// SetFromMap updates settings.
func (s *Settings) SetFromMap(ctx context.Context, settingsMap map[string]any) gameerr.GameErrorer {
	if len(settingsMap) == 0 {
		return nil
	}

	data, err := json.Marshal(settingsMap)
	if err != nil {
		return gameerr.New(
			gameerr.ErrCodeSettingsMarshal,
			"failed to marshal input map",
			err,
			nil,
		)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	err = json.Unmarshal(data, s)

	if err != nil {
		return gameerr.New(
			gameerr.ErrCodeSettingsUnmarshal,
			"failed to unmarshal settings from map (type mismatch?)",
			err,
			map[string]any{
				"input": settingsMap,
			},
		).LogAndReturn(ctx, slog.LevelError)
	}

	return nil
}

// ToMap returns the current settings as a map.
func (s *Settings) ToMap(ctx context.Context) (map[string]any, gameerr.GameErrorer) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	data, err := json.Marshal(s)
	if err != nil {
		return nil, gameerr.New(
			gameerr.ErrCodeSettingsMarshal,
			"failed to marshal input map",
			err,
			nil,
		).LogAndReturn(ctx, slog.LevelError)
	}

	var result map[string]any

	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, gameerr.New(
			gameerr.ErrCodeSettingsUnmarshal,
			"failed to unmarshal settings from map (type mismatch?)",
			err,
			nil,
		).LogAndReturn(ctx, slog.LevelError)
	}

	return result, nil
}

// GetSingle returns a single value by key name.
func (s *Settings) GetSingle(ctx context.Context, key SettingKey) (any, bool) {
	allSettings, err := s.ToMap(ctx)
	if err != nil {
		return nil, false
	}
	value, ok := allSettings[string(key)]
	return value, ok
}

func NewSettings(ctx context.Context) (*Settings, gameerr.GameErrorer) {
	file, err := os.Open(settingsFileName)

	if err != nil {
		gameerr.New(
			gameerr.ErrCodeSettingsLoading,
			"settings file is missing",
			err,
			map[string]any{
				"filename": settingsFileName,
				"attempt":  1,
			},
		).Log(ctx, slog.LevelWarn)
		return createDefaultSettings(), nil
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(gameerr.New(
				gameerr.ErrCodeSettingsClosing,
				"settings file closing failed",
				err,
				nil,
			).LogAndReturn(ctx, slog.LevelError))
		}
	}(file)

	var s Settings

	if err := json.NewDecoder(file).Decode(&s); err != nil {
		gameerr.New(
			gameerr.ErrCodeSettingsReading,
			"settings file is corrupted",
			err,
			map[string]any{
				"filename": settingsFileName,
			},
		).Log(ctx, slog.LevelError)
		return createDefaultSettings(), nil
	}

	return &s, nil
}

func createDefaultSettings() *Settings {
	return &Settings{
		PlayerName:  fmt.Sprintf("Player%d", rand.Intn(100)),
		PlayerColor: "DC143C",
		Locale:      "en_US",
	}
}
