package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	game "gitlab.utc.fr/vilecler/musicslime/game"
	types "gitlab.utc.fr/vilecler/musicslime/types"
)

func main() {
	fmt.Println("Start of the program")

	ebiten.SetWindowSize(int(types.GetWindowDefault().Width), int(types.GetWindowDefault().Height))
	ebiten.SetWindowTitle("Music Slime Demo")
	if err := ebiten.RunGame(game.NewGame()); err != nil {
		log.Fatal(err)
	}
}
