package audio_system

import (
	"github.com/KelleyTyler/ebiten_basic_framework/internal/audio_system"
	"github.com/KelleyTyler/ebiten_basic_framework/internal/settings"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

type Audio_System struct {
	Audio_Context *audio.Context
	Settings      *settings.Game_Settings
	A_System      audio_system.Audio_System
	Sounds        [][]byte
}
