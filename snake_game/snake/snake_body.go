package snake

import (
	"fyne.io/fyne/v2"
	"time"
)

const (
	CellSize   = 20
	GridWidth  = 30
	GridHeight = 20
)

var (
	Directions = []fyne.KeyName{
		fyne.KeyUp,
		fyne.KeyDown,
		fyne.KeyLeft,
		fyne.KeyRight,
	}
)

type Snake struct {
	Body       []fyne.Position
	Length     int
	Direction  fyne.KeyName
	lastMoveAt time.Time
}

func NewSnake() Snake {
	head := fyne.NewPos(CellSize*GridWidth/2, CellSize*GridHeight/2)
	return Snake{
		Body:       []fyne.Position{head},
		Length:     1,
		Direction:  fyne.KeyRight,
		lastMoveAt: time.Now(),
	}
}

func (s *Snake) Move() {
	if time.Since(s.lastMoveAt) < 100*time.Millisecond {
		return
	}

	head := s.Body[0]
	var newHead fyne.Position

	switch s.Direction {
	case fyne.KeyUp:
		newHead = fyne.NewPos(head.X, head.Y-CellSize)
	case fyne.KeyDown:
		newHead = fyne.NewPos(head.X, head.Y+CellSize)
	case fyne.KeyLeft:
		newHead = fyne.NewPos(head.X-CellSize, head.Y)
	case fyne.KeyRight:
		newHead = fyne.NewPos(head.X+CellSize, head.Y)
	}

	s.Body = append([]fyne.Position{newHead}, s.Body[:s.Length-1]...)
	s.lastMoveAt = time.Now()
}

func (s *Snake) EatFood() {
	s.Length++
	s.Body = append([]fyne.Position{s.Body[0]}, s.Body...)
}

func (s *Snake) CollideWithSelf() bool {
	head := s.Body[0]
	for _, b := range s.Body[1:] {
		if head == b {
			return true
		}
	}
	return false
}

func (s *Snake) CollideWithWall() bool {
	head := s.Body[0]
	return head.X < 0 || head.X >= CellSize*GridWidth || head.Y < 0 || head.Y >= CellSize*GridHeight
}
