// game/board.go
package game

import (
	"math/rand"
	"time"
)

type GameBoard struct {
	Width, Height int
	Snake         *Snake
	Foods         []Apple
	Walls         []Wall
	Timer         time.Time
}

func NewGameBoard(width, height int) *GameBoard {
	snake := NewSnake(Position{X: width / 2, Y: height / 2})
	walls := generateWalls(width, height, 10, snake)

	return &GameBoard{
		Width:  width,
		Height: height,
		Snake:  snake,
		Walls:  walls,
		Foods:  generateFood(width, height, 2, snake, walls),
		Timer:  time.Now(),
	}
}

func (g *GameBoard) Update() {
	if g.Snake.Alive {
		g.Snake.ProcessInput()
		g.Snake.LimitSpeed(10000)

		if time.Since(g.Timer).Milliseconds() >= g.Snake.MoveInterval() {
			g.Snake.Move()
			g.Timer = time.Now()
		}

		g.checkFoodCollision()
		g.checkWallCollision()
		g.checkSelfCollision()
	}
}

func (g *GameBoard) checkFoodCollision() {
	for i := len(g.Foods) - 1; i >= 0; i-- {
		if g.Snake.Body[0].Equals(g.Foods[i].Position()) {
			g.Foods[i].Eat(g.Snake)
			g.Foods = append(g.Foods[:i], g.Foods[i+1:]...)
			g.Foods = append(g.Foods, generateFood(g.Width, g.Height, 1, g.Snake, g.Walls)...)
		}
	}
}

func (g *GameBoard) checkWallCollision() {
	for i := len(g.Walls) - 1; i >= 0; i-- {
		if g.Snake.Body[0].Equals(g.Walls[i].Position()) {
			g.Walls[i].Interact(g.Snake)

			if _, isSoft := g.Walls[i].(*SoftWall); isSoft {
				g.Walls = append(g.Walls[:i], g.Walls[i+1:]...)
			}
		}
	}
}

func (g *GameBoard) checkSelfCollision() {
	for i := 1; i < len(g.Snake.Body)-1; i++ {
		if g.Snake.Body[0].Equals(g.Snake.Body[i]) {
			g.Snake.Alive = false
			break
		}
	}
}

func generateFood(width, height, foodCount int, snake *Snake, walls []Wall) []Apple {
	var foodItems []Apple

	for len(foodItems) < foodCount {
		x, y := randomPosition(width, height)
		newFood := NewNormalApple(x, y)

		if !isPositionOccupied(newFood.Position(), snake, walls) {
			if rand.Intn(5) == 0 {
				foodItems = append(foodItems, NewSlowApple(x, y))
			} else {
				foodItems = append(foodItems, newFood)
			}
		}
	}

	return foodItems
}

func generateWalls(width, height, wallCount int, snake *Snake) []Wall {
	var walls []Wall

	for len(walls) < wallCount {
		x, y := randomPosition(width, height)
		newWall := NewSoftWall(x, y)

		if !isPositionOccupied(newWall.Position(), snake, walls) {
			if rand.Intn(2) == 0 {
				walls = append(walls, NewHardWall(x, y))
			} else {
				walls = append(walls, NewSoftWall(x, y))
			}
		}
	}

	return walls
}

func randomPosition(width, height int) (int, int) {
	x := rand.Intn(width)
	y := rand.Intn(height)
	return x, y
}

func isPositionOccupied(pos Position, snake *Snake, walls []Wall) bool {
	for _, bodyPart := range snake.Body {
		if bodyPart.Equals(pos) {
			return true
		}
	}
	for _, wall := range walls {
		if wall.Position().Equals(pos) {
			return true
		}
	}

	return false
}
