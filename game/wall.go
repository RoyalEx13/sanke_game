// game/wall.go
package game

import "image/color"

type Wall interface {
	Interact(snake *Snake)
	Position() Position
	Color() color.Color
}

type BaseWall struct {
	pos   Position
	color color.Color
}

func (w *BaseWall) Position() Position {
	return w.pos
}

func (w *BaseWall) Color() color.Color {
	return w.color
}

func NewBaseWall(x, y int, col color.Color) *BaseWall {
	return &BaseWall{pos: Position{X: x, Y: y}, color: col}
}

type SoftWall struct {
	*BaseWall
}

func NewSoftWall(x, y int) *SoftWall {
	return &SoftWall{BaseWall: NewBaseWall(x, y, color.RGBA{0, 128, 255, 255})}
}

func (w *SoftWall) Interact(snake *Snake) {
	snake.ReduceSize()
}

type HardWall struct {
	*BaseWall
}

func NewHardWall(x, y int) *HardWall {
	return &HardWall{BaseWall: NewBaseWall(x, y, color.RGBA{0, 0, 255, 255})}
}

func (w *HardWall) Interact(snake *Snake) {
	snake.Die()
}
