package user_interface

import (
	"bytes"
	"fmt"
	"log"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/font/gofont/goregular"

	"github.com/KelleyTyler/ebiten_basic_framework/internal/audio_system"
	"github.com/KelleyTyler/ebiten_basic_framework/internal/settings"
)

/* the things that are needed here;
 * -- to have a framework
 *
 *
 *
 *
 *
 */

type UI_Backend struct {
	Settings       *settings.Game_Settings
	E_UI           *ebitenui.UI
	AudioSystem    *audio_system.Audio_System
	Text_Faces_Src []*text.GoTextFaceSource
	UI_Sounds      [][]byte
	isinit         bool
}

func Get_UI_Backend(GSetting *settings.Game_Settings, aud_sys *audio_system.Audio_System) UI_Backend {

	bckend := UI_Backend{
		Settings:    GSetting,
		AudioSystem: aud_sys,
	}
	bckend.Text_Faces_Src = make([]*text.GoTextFaceSource, 0)
	var err error = nil
	var tempTextSrc *text.GoTextFaceSource
	tempTextSrc, err = text.NewGoTextFaceSource(bytes.NewReader(gomono.TTF))
	if err != nil {
		log.Fatal("err: ", err)
	}
	bckend.Text_Faces_Src = append(bckend.Text_Faces_Src, tempTextSrc)

	tempTextSrc, err = text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal("err: ", err)
	}
	bckend.Text_Faces_Src = append(bckend.Text_Faces_Src, tempTextSrc)

	bckend.Init_UI_Sounds()
	root := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	bckend.E_UI = &ebitenui.UI{
		Container: root,
	}

	bckend.isinit = true
	return bckend
}

func (ui_bckend *UI_Backend) Update() (err error) {
	ui_bckend.E_UI.Update()

	return err
}

func (ui_bckend *UI_Backend) Draw(screen *ebiten.Image) (err error) {
	ui_bckend.E_UI.Draw(screen)

	return err
}

/*
 */
func (ui_bckend *UI_Backend) Init_UI_Sounds() {
	ui_bckend.UI_Sounds = append(ui_bckend.UI_Sounds, audio_system.Soundwave_CreateSound(3200, 200, 0, 110, []float32{1.0}, []float32{0.0750000}))
	ui_bckend.UI_Sounds = append(ui_bckend.UI_Sounds, audio_system.Soundwave_CreateSound(3200, 200, 10, 110, []float32{1.0}, []float32{0.0750000}))
	ui_bckend.UI_Sounds = append(ui_bckend.UI_Sounds, audio_system.Soundwave_CreateSound(3200, 200, 15, 110, []float32{1.0}, []float32{0.0750000}))
	ui_bckend.UI_Sounds = append(ui_bckend.UI_Sounds, audio_system.Soundwave_CreateSound(3200, 200, 20, 110, []float32{1.0}, []float32{0.0750000}))
	ui_bckend.UI_Sounds = append(ui_bckend.UI_Sounds, audio_system.Soundwave_CreateSound(3200, 200, 25, 110, []float32{1.0}, []float32{0.0750000}))
	ui_bckend.UI_Sounds = append(ui_bckend.UI_Sounds, audio_system.Soundwave_CreateSound(3200, 200, 25, 110, []float32{1.0}, []float32{0.0750000}))
}

/**/
func (ui_bckend *UI_Backend) Play_UI_Sound(num int) (err error) {
	err = nil
	if num < int(len(ui_bckend.UI_Sounds)) {
		ui_bckend.AudioSystem.Play_Sound_FromBytes(ui_bckend.UI_Sounds[num])
	} else {
		return fmt.Errorf("tried to play sound that does not exist")
	}
	return err
}

/*
 */
func (ui_bckend *UI_Backend) Play_Byte_Sound(bytes []byte) {
	ui_bckend.AudioSystem.Play_Sound_FromBytes(bytes)
}

// /*
//  */
// func (ui_bckend *UI_Backend) Create_Own_Sound_System() (err error) {
// 	ui_bckend.AudioSystem=
// 	return err
// }
