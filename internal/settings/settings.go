package settings

/*
 *
 *
 */

type Game_Settings struct {
	Version       string
	Window_Size_X uint
	Window_Size_Y uint
	Window_Res_X  uint
	Window_Res_Y  uint

	UI_Volume uint8
}

func Get_Basic_Settings() Game_Settings {
	gSets := Game_Settings{
		Version:       "0.0.0",
		Window_Size_X: 960,
		Window_Size_Y: 640,
		Window_Res_X:  960,
		Window_Res_Y:  640,
		UI_Volume:     255,
	}
	return gSets
}
