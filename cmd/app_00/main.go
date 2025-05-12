package main

import (
	"log"

	"github.com/KelleyTyler/ebiten_basic_framework/internal/game_framework"
	"github.com/hajimehoshi/ebiten/v2"
)

// type Game struct{}

// func (g *Game) Update() error {
// 	return nil
// }

// func (g *Game) Draw(screen *ebiten.Image) {
// 	ebitenutil.DebugPrint(screen, "Hello, World!")
// }

// func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
// 	return 320, 240
// }

func main() {
	my_game := game_framework.Game{}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&my_game); err != nil {
		log.Fatal(err)
	}
}
