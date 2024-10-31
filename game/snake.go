// game/snake.go
package game

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Snake struct {
	Body          []Position
	Direction     Direction
	LastDirection Direction
	Speed         int64
	Alive         bool
	IsBoosted     bool
}

func NewSnake(start Position) *Snake {
	return &Snake{
		Body:      []Position{start, {start.X + 1, start.Y}, {start.X + 2, start.Y}},
		Direction: Right,
		Speed:     100,
		Alive:     true,
		IsBoosted: false,
	}
}

func (s *Snake) Move() {
	head := s.Body[0]
	newHead := head

	switch s.Direction {
	case Up:
		newHead.Y--
		s.LastDirection = Up
	case Down:
		newHead.Y++
		s.LastDirection = Down
	case Left:
		newHead.X--
		s.LastDirection = Left
	case Right:
		newHead.X++
		s.LastDirection = Right
	}

	if newHead.X < 0 {
		newHead.X = 39
	} else if newHead.X >= 40 {
		newHead.X = 0
	}
	if newHead.Y < 0 {
		newHead.Y = 39
	} else if newHead.Y >= 40 {
		newHead.Y = 0
	}

	s.Body = append([]Position{newHead}, s.Body[:len(s.Body)-1]...)
}

func (s *Snake) Grow() {
	s.Body = append(s.Body, s.Body[len(s.Body)-1])
}

func (s *Snake) SpeedUp() {
	s.Speed += 10
}

func (s *Snake) SlowDown() {
	if s.Speed > 1 {
		s.Speed -= 10
	} else {
		s.Speed = 1
	}
}

func (s *Snake) ReduceSize() {
	if len(s.Body) > 1 {
		s.Body = s.Body[:len(s.Body)-1]
	} else {
		s.Die()
	}
}

func (s *Snake) Die() {
	s.Alive = false
}

func (s *Snake) LimitSpeed(max int64) {
	if s.Speed > max {
		s.Speed = max
	}
}

func (s *Snake) MoveInterval() int64 {
	if s.IsBoosted {
		return 10000 / (s.Speed * 2)
	} else {
		return 10000 / s.Speed
	}
}

func (s *Snake) BootSpeed() int64 {
	return s.Speed * 2
}
