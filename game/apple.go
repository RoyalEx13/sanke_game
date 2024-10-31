// game/apple.go
package game

import "image/color"

type Apple interface {
	Eat(snake *Snake)
	Position() Position
	Color() color.Color
}

type BaseApple struct {
	pos   Position
	color color.Color
}

func NewBaseApple(x, y int, col color.Color) *BaseApple {
	return &BaseApple{pos: Position{X: x, Y: y}, color: col}
}

func (a *BaseApple) Position() Position {
	return a.pos
}

func (a *BaseApple) Color() color.Color {
	return a.color
}

type NormalApple struct {
	*BaseApple
}

func NewNormalApple(x, y int) *NormalApple {
	return &NormalApple{BaseApple: NewBaseApple(x, y, color.RGBA{255, 0, 0, 255})}
}

func (a *NormalApple) Eat(snake *Snake) {
	snake.Grow()
	snake.SpeedUp()
}

type SlowApple struct {
	*BaseApple
}

func NewSlowApple(x, y int) *SlowApple {
	return &SlowApple{BaseApple: NewBaseApple(x, y, color.RGBA{90, 34, 139, 255})}
}

func (a *SlowApple) Eat(snake *Snake) {
	snake.Grow()
	snake.SlowDown()
}
