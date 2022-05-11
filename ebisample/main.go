package main

import (
	"log"

	"github.com/belong-try/ebisample/tictactoe"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game, err := tictactoe.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(960, 960)
	ebiten.SetWindowTitle("Tic Tac Toe")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
