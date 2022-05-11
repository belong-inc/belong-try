package tictactoe

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Turn int

const (
	Circle Turn = 0
	Cross  Turn = 1
)

type Game struct {
	images *Images

	turn Turn
}

type Images struct {
	game   *ebiten.Image
	board  *ebiten.Image
	circle *ebiten.Image
	cross  *ebiten.Image
}

func NewGame() (*Game, error) {
	circleImage, _, err := ebitenutil.NewImageFromFile("images/circle.png")
	if err != nil {
		return nil, err
	}
	crossImage, _, err := ebitenutil.NewImageFromFile("images/cross.png")
	if err != nil {
		return nil, err
	}
	boardImage, _, err := ebitenutil.NewImageFromFile("images/board.png")
	if err != nil {
		return nil, err
	}
	game := &Game{
		images: &Images{
			game:   ebiten.NewImage(480, 480),
			board:  boardImage,
			circle: circleImage,
			cross:  crossImage,
		},
	}
	return game, nil
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(x-60), float64(y-60))
		switch g.turn {
		case Circle:
			g.images.board.DrawImage(g.images.circle, opts)
		case Cross:
			g.images.board.DrawImage(g.images.cross, opts)
		}
		g.turn ^= 1
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.images.board, nil)
	screen.DrawImage(g.images.game, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 480, 480
}
