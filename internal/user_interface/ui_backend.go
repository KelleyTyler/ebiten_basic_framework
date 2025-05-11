package user_interface

import (
	"github.com/ebitenui/ebitenui"

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
	Settings *settings.Game_Settings
	E_UI     *ebitenui.UI
}
