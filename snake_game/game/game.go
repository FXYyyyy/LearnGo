package game

import (
	"firstGo/snake_game/snake"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"math/rand"
)

type Game struct {
	Snake   snake.Snake
	Food    fyne.Position
	Score   int
	Over    bool
	WinSize fyne.Size
}

func NewGame(winSize fyne.Size) *Game {
	game := &Game{
		Snake:   snake.NewSnake(),
		Score:   0,
		Over:    false,
		WinSize: winSize,
	}
	game.SetRandFood()
	return game
}

func (g *Game) Update() {
	if g.Over {
		return
	}

	g.Snake.Move()

	if g.Snake.CollideWithSelf() || g.Snake.CollideWithWall() {
		g.Over = true
		return
	}

	if g.Snake.Body[0] == g.Food {
		g.Snake.EatFood()
		g.SetRandFood()
		g.Score++
	}
}

func (g *Game) Draw() fyne.CanvasObject {
	snakeColor := color.RGBA{0, 255, 0, 255}
	foodColor := color.RGBA{255, 0, 0, 255}

	objects := make([]fyne.CanvasObject, 0)

	for _, b := range g.Snake.Body {
		temp := canvas.NewRectangle(snakeColor)
		temp.SetMinSize(fyne.NewSize(snake.CellSize, snake.CellSize))
		temp.Move(b)
		objects = append(objects, temp)
	}

	foodObject := canvas.NewCircle(foodColor)
	foodObject.Move(g.Food)
	objects = append(objects, foodObject)

	scoreLabel := widget.NewLabel(fmt.Sprintf("Score: %d", g.Score))
	scoreLabel.Move(fyne.NewPos(g.WinSize.Width-scoreLabel.MinSize().Width-10, 10))
	objects = append(objects, scoreLabel)

	return container.NewWithoutLayout(objects...)
}

func (g *Game) SetDirection(direction fyne.KeyName) {
	if direction == fyne.KeyReturn && g.Over {
		*g = *NewGame(g.WinSize)
		return
	}
	if contains(snake.Directions, direction) {
		g.Snake.Direction = direction
	}
}

func (g *Game) SetRandFood() {
	g.Food = fyne.NewPos(float32(rand.Intn(snake.GridWidth)*snake.CellSize), float32(rand.Intn(snake.GridHeight)*snake.CellSize))
	return
}

func contains(keys []fyne.KeyName, key fyne.KeyName) bool {
	for _, k := range keys {
		if k == key {
			return true
		}
	}
	return false
}
