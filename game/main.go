package game

import (
	"fmt"
	"time"

	"github.com/arielril/go-space-invaders/util"
	"github.com/go-gl/gl/v2.1/gl"
	glfw "github.com/go-gl/glfw/v3.3/glfw"
)

var gameShips []Ship
var gameCar Car
var gameBullets []Bullet
var fps util.FPS
var maxFps float64 = 30
var startTime time.Time

const (
	gameScale = .1
	maxShoots = 15
)

// Init the game objects
func Init() {
	gameShips = []Ship{
		NewShip(ships[Ship2], Ship2).SetPos(5, 5).(Ship),
		// NewShip(ships[Ship3], Ship3).SetPos(2, 10).(Ship),
		// NewShip(ships[Ship1], Ship1).SetPos(7, 10).(Ship),
		// NewShip(ships[Ship4], Ship4).SetPos(0, 10).(Ship),
	}

	gameCar = NewCar(carData).SetPos(10, 0).(Car)
	playerLives = []Life{
		NewLife(lifeData, 17, 9),
		NewLife(lifeData, 18, 9),
		NewLife(lifeData, 19, 9),
	}
	fps = util.NewFps()
	startTime = time.Now()
	fmt.Printf("Started at: %v\n", startTime)
}

func drawShips() {
	for _, v := range gameShips {
		v.SetScale(gameScale).Draw()
		// v.SetSpeed(float32(fps.GetFPS())).Move()
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

func drawLives() {
	for _, l := range playerLives {
		l.SetScale(gameScale).Draw()
	}
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

func drawBoundingBoxes() {
	for _, b := range gameBullets {
		if b == nil {
			continue
		}
		b.GetBoundingBox().Draw()
	}

	for _, s := range gameShips {
		if s == nil {
			continue
		}
		s.GetBoundingBox().Draw()
	}
}

func displayFps() {
	acc := fps.SetFPS().GetAccumulated()
	if acc >= 1 { // print every second
		fmt.Printf("FPS: %v\n", fps.GetFPS())
		fps.Reset()
	}
}

// Display the game
func Display(w *glfw.Window) {
	startTime := time.Now()
	displayFps()

	// draw objects
	drawGround()
	drawLives()
	drawCar()
	// drawBoundingBoxes()
	drawShips()
	drawBullets()

	// check for collisions
	doCollisions()

	// clear game objects/screen
	removeBulletsFromGame(w)
	optimizeGame()

	if !IsAlive() {
		// TODO end game
	}

	time.Sleep(
		time.Second/time.Duration(maxFps) - time.Since(startTime),
	)
}
