// game/input.go
package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (s *Snake) ProcessInput() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && s.LastDirection != Down {
		s.Direction = Up
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && s.LastDirection != Up {
		s.Direction = Down
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && s.LastDirection != Right {
		s.Direction = Left
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && s.LastDirection != Left {
		s.Direction = Right
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		s.IsBoosted = true
	} else {
		s.IsBoosted = false
	}

}
