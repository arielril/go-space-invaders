package game

import (
	"github.com/go-gl/gl/v2.1/gl"
	glfw "github.com/go-gl/glfw/v3.3/glfw"
)

var gameShips []Ship
var gameCar Car
var gameBullets []Bullet

const (
	gameScale = .1
	maxShoots = 15
)

// GetCar returns the car object
func GetCar() Car {
	return gameCar
}

// AddShoot add a new bullet to the game
func AddShoot(b Bullet) {
	if len(gameBullets) < maxShoots {
		gameBullets = append(gameBullets, b)
	}
}

// Init the game objects
func Init() {
	gameShips = append(
		gameShips,
		NewShip(ships[Ship1], Ship1).SetPos(-8, 9).(Ship),
		NewShip(ships[Ship2], Ship2).SetPos(-5, 9).(Ship),
		NewShip(ships[Ship1], Ship1).SetPos(-3, 9).(Ship),
		NewShip(ships[Ship1], Ship1).SetPos(0, 9).(Ship),
		NewShip(ships[Ship4], Ship4).SetPos(3, 9).(Ship),
		NewShip(ships[Ship1], Ship1).SetPos(5, 9).(Ship),
		NewShip(ships[Ship3], Ship3).SetPos(8, 9).(Ship),

		NewShip(ships[Ship3], Ship3).SetPos(-8, 7).(Ship),
		NewShip(ships[Ship3], Ship3).SetPos(-5, 7).(Ship),
		NewShip(ships[Ship3], Ship3).SetPos(-3, 7).(Ship),
		NewShip(ships[Ship3], Ship3).SetPos(0, 7).(Ship),
		NewShip(ships[Ship3], Ship3).SetPos(3, 7).(Ship),
		NewShip(ships[Ship3], Ship3).SetPos(5, 7).(Ship),
		NewShip(ships[Ship3], Ship3).SetPos(8, 7).(Ship),

		NewShip(ships[Ship1], Ship1).SetPos(-8, 5).(Ship),
		NewShip(ships[Ship2], Ship2).SetPos(-5, 5).(Ship),
		NewShip(ships[Ship1], Ship1).SetPos(-3, 5).(Ship),
		NewShip(ships[Ship1], Ship1).SetPos(0, 5).(Ship),
		NewShip(ships[Ship2], Ship2).SetPos(3, 5).(Ship),
		NewShip(ships[Ship1], Ship1).SetPos(5, 5).(Ship),
		NewShip(ships[Ship4], Ship4).SetPos(8, 5).(Ship),
	)

	gameCar = NewCar(carData).SetPos(0, 0).(Car)
}

func drawShips() {
	for _, v := range gameShips {
		v.SetScale(gameScale).Draw()
	}
}

func drawBullets() {
	for _, v := range gameBullets {
		if v == nil {
			continue
		}
		v.SetScale(gameScale).Draw()
		v.Move()
	}
}

func drawCar() {
	gameCar.SetScale(gameScale).Draw()
}

func drawGround() {
	gl.PushMatrix()
	{
		gl.Color3f(.08, .5, 0) // #158000
		gl.LineWidth(10)

		gl.Begin(gl.LINES)
		{
			gl.Vertex2f(-20, .020)
			gl.Vertex2f(20, .020)
		}
		gl.End()
	}
	gl.PopMatrix()
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

func checkShipKill() {
	for _, bull := range gameBullets {
		if bull == nil {
			continue
		}

		for _, sh := range gameShips {
			bull.CheckHit(sh)
			// if the bullet hit the ship, the ship must die :D
		}

	}
}

// Display the game
func Display(w *glfw.Window) {
	// draw objects
	drawGround()
	drawCar()
	drawShips()
	drawBullets()

	// TODO: check hit
	// TODO: kill ships
	// TODO: add lives for the player

	// clear game objects/screen
	removeBulletsFromGame(w)
	optimizeGame()
}
