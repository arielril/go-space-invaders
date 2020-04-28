package game

import (
	"github.com/go-gl/gl/v2.1/gl"
	glfw "github.com/go-gl/glfw/v3.3/glfw"
)

var gameShips []Ship
var gameCar Car
var gameBullets []Bullet

const gameScale = .1

// Init the game objects
func Init() {
	gameShips = append(
		gameShips,
		NewShip(ships[Ship1], Ship1).SetPos(-8, 9),
		NewShip(ships[Ship1], Ship1).SetPos(-5, 9),
		NewShip(ships[Ship1], Ship1).SetPos(-3, 9),
		NewShip(ships[Ship1], Ship1).SetPos(0, 9),
		NewShip(ships[Ship1], Ship1).SetPos(3, 9),
		NewShip(ships[Ship1], Ship1).SetPos(5, 9),
		NewShip(ships[Ship1], Ship1).SetPos(8, 9),

		NewShip(ships[Ship1], Ship1).SetPos(-8, 7),
		NewShip(ships[Ship1], Ship1).SetPos(-5, 7),
		NewShip(ships[Ship1], Ship1).SetPos(-3, 7),
		NewShip(ships[Ship1], Ship1).SetPos(0, 7),
		NewShip(ships[Ship1], Ship1).SetPos(3, 7),
		NewShip(ships[Ship1], Ship1).SetPos(5, 7),
		NewShip(ships[Ship1], Ship1).SetPos(8, 7),

		NewShip(ships[Ship1], Ship1).SetPos(-8, 5),
		NewShip(ships[Ship1], Ship1).SetPos(-5, 5),
		NewShip(ships[Ship1], Ship1).SetPos(-3, 5),
		NewShip(ships[Ship1], Ship1).SetPos(0, 5),
		NewShip(ships[Ship1], Ship1).SetPos(3, 5),
		NewShip(ships[Ship1], Ship1).SetPos(5, 5),
		NewShip(ships[Ship1], Ship1).SetPos(8, 5),
	)
}

func drawShips() {
	for _, v := range gameShips {
		v.SetScale(gameScale).Draw()
	}
}

func drawBullets() {
	for _, v := range gameBullets {
		v.SetScale(gameScale).Draw()
	}
}

func drawCar() {
	gameCar.SetScale(gameScale).Draw()
}

func showAxis() {
	gl.PushMatrix()
	{
		gl.Color3f(0, 0, 0)
		gl.LineWidth(5)

		gl.Begin(gl.LINES)

		gl.Vertex2f(-20, 0)
		gl.Vertex2f(20, 0)

		gl.Vertex2f(0, 20)
		gl.Vertex2f(0, -20)

		gl.End()
	}
	gl.PopMatrix()
}

// Display the game
func Display(w *glfw.Window) {
	showAxis()

	drawShips()
}
