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

func (t Turn) String() string {
	switch t {
	case Circle:
		return "○"
	case Cross:
		return "×"
	default:
		return ""
	}
}

type Game struct {
	images *Images
	board  [3][3]string
	turn   Turn
	count  int
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
		// cool code.
		// opts.GeoM.Translate(float64(160*(x/160+1)-160), float64(160*(y/160+1)-160))
		opts.GeoM.Translate(getBoardPosition(x), getBoardPosition(y))
		if !g.isEmpty(x, y) {
			return nil
		}
		g.setInternalBoardPosition(x, y)
		switch g.turn {
		case Circle:
			g.images.board.DrawImage(g.images.circle, opts)
		case Cross:
			g.images.board.DrawImage(g.images.cross, opts)
		}
		g.count++
		if g.getWinner() != "" {
			return nil
		}
		if g.checkGameEnd() {
			g.images.board.Clear()
			return nil
		}
		g.turn ^= 1
	}

	return nil
}

func (g *Game) isEmpty(x, y int) bool {
	return g.board[x/160][y/160] == ""
}

func (g *Game) setInternalBoardPosition(x, y int) {
	g.board[x/160][y/160] = g.turn.String()
}

func (g *Game) checkGameEnd() bool {
	return g.count == 9
}

func (g *Game) getWinner() string {
	for i := range g.board {
		if isEqual(g.board[i][0], g.board[i][1], g.board[i][2]) {
			return g.turn.String()
		}
		if isEqual(g.board[0][i], g.board[1][i], g.board[2][i]) {
			return g.turn.String()
		}
	}
	if isEqual(g.board[0][0], g.board[1][1], g.board[2][2]) {
		return g.turn.String()
	}
	if isEqual(g.board[0][2], g.board[1][1], g.board[2][0]) {
		return g.turn.String()
	}
	return ""
}

func isEqual(s1, s2, s3 string) bool {
	return s1 != "" && s1 == s2 && s2 == s3
}

func getBoardPosition(pos int) float64 {
	if 320 < pos {
		return 320
	}
	if 160 < pos {
		return 160
	}
	return 0
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.images.board, nil)
	screen.DrawImage(g.images.game, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 480, 480
}
