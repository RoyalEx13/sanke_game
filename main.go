// main.go
package main

import (
	"image/color"
	"log"

	"github.com/RoyalEx13/sanke_game/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type SnakeGame struct {
	board      *game.GameBoard
	appleImage *ebiten.Image
}

func NewSnakeGame() *SnakeGame {
	return &SnakeGame{
		board: game.NewGameBoard(40, 40),
	}
}

func (g *SnakeGame) Reset() {
	g.board = game.NewGameBoard(40, 40)
}

func (g *SnakeGame) Update() error {
	if g.board.Snake.Alive {
		g.board.Snake.ProcessInput()
		g.board.Update()
	}
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		g.Reset()
	}
	return nil
}

func (g *SnakeGame) Draw(screen *ebiten.Image) {
	for _, part := range g.board.Snake.Body {
		vector.DrawFilledRect(screen, float32(part.X*40), float32(part.Y*40), 36, 36, color.RGBA{0, 255, 0, 255}, false)
	}
	for _, food := range g.board.Foods {
		vector.DrawFilledCircle(screen, float32(food.Position().X*40)+18, float32(food.Position().Y*40)+18, 18, food.Color(), false)
	}

	for _, wall := range g.board.Walls {
		vector.DrawFilledRect(screen, float32(wall.Position().X*40), float32(wall.Position().Y*40), 36, 36, wall.Color(), false)
	}
}

func (g *SnakeGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.board.Width * 40, g.board.Height * 40
}

func main() {
	ebiten.SetWindowSize(600, 600)
	ebiten.SetWindowTitle("Snake Game")
	if err := ebiten.RunGame(NewSnakeGame()); err != nil {
		log.Fatal(err)
	}
}
