package game

import (
	"github.com/go-gl/gl/v2.1/gl"
	glfw "github.com/go-gl/glfw/v3.3/glfw"
)

func drawLines() {
	gl.Begin(gl.LINES)
	gl.Vertex2f(0, 0)
	gl.Vertex2f(5, 5)
	gl.End()

	gl.Begin(gl.LINES)
	gl.Vertex2f(5, 5)
	gl.Vertex2f(10, 0)
	gl.End()
}

func siren() {
	gl.Begin(gl.LINES)
	{
		gl.Vertex2f(-1, -1)
		gl.Vertex2f(1, 1)
		gl.Vertex2f(1, -1)
		gl.Vertex2f(-1, 1)
	}
	gl.End()
}

func smallSiren() {
	gl.PushMatrix()
	{
		gl.Scalef(.3, .3, 1)
		siren()
	}
	gl.PopMatrix()
}

func spinningSmallSiren(ang float32) {
	gl.PushMatrix()
	{
		gl.Rotatef(ang, 0, 0, 1)
		smallSiren()
	}
	gl.PopMatrix()
}

func drawCarChassis() {
	gl.Begin(gl.QUADS)
	{
		gl.Vertex2f(0, 0)
		gl.Vertex2f(-2, 0)
		gl.Vertex2f(-2, 1)
		gl.Vertex2f(0, 1)
	}
	gl.End()
}

var sirenAngle float32

func drawCar() {
	gl.PushMatrix()
	{
		drawCarChassis()
		gl.Color3f(0, 1, 0)
		gl.Translatef(-1.5, 1.3, 0)
		spinningSmallSiren(sirenAngle)
		sirenAngle++
	}
	gl.PopMatrix()
}

var x float32

// Display the game
func Display(w *glfw.Window) {
	x += .01

	gl.Color3f(1, 0, 0)
	gl.Translatef(5, 5, 0)

	drawCar()

	gl.Color3f(0, 0, 1)
	gl.Translatef(x, 0, 0)
	drawCar()
}
