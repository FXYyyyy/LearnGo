package main

import (
	"firstGo/snake_game/game"
	"firstGo/snake_game/snake"
	"fyne.io/fyne/v2"
	app "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"image"
	"image/color"
	"image/draw"
)

func main() {
	myGame := game.NewGame(fyne.NewSize(snake.GridWidth*snake.CellSize, snake.GridHeight*snake.CellSize))

	app := app.New()
	window := app.NewWindow("Snake Game")
	window.Resize(myGame.WinSize)

	window.SetContent(canvas.NewRasterWithPixels(func(x, y, w, h int) color.Color {
		img := image.NewRGBA(image.Rect(x, y, w, h))
		draw.Draw(img, img.Bounds(), &image.Uniform{C: color.Black}, image.ZP, draw.Src)

		canvasObj := myGame.Draw()
		canvasObj.Show()

		return
	}))
}
