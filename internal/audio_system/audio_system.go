package audio_system

import (
	// "github.com/KelleyTyler/ebiten_basic_framework/internal/audio_system"
	"github.com/KelleyTyler/ebiten_basic_framework/internal/settings"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

/*
This is designed to be a simple system
*/
type Audio_System struct {
	Audio_Context *audio.Context
	Settings      *settings.Game_Settings
	Base_Freq     int
	Sameple_Rate  int
	// Sounds        [][]byte
}

/*
when a byte array is inputted this will play it such that
*/
func (aud *Audio_System) Play_Sound_FromBytes(bytes []byte) {
	p := aud.Audio_Context.NewPlayerF32FromBytes(bytes)
	p.SetVolume(float64(aud.Settings.UI_Volume) / 100)
	p.Play()
}

/*
returns an audio-system;
*/
func Get_Audio_System(GS *settings.Game_Settings, sRate, bFreq int) (audio_sys Audio_System) {
	audio_sys.Base_Freq = bFreq
	audio_sys.Sameple_Rate = sRate
	audio_sys.Settings = GS
	audio_sys.Audio_Context = audio.NewContext(sRate)
	return audio_sys
}
