package game

import (
	"github.com/go-gl/gl/v2.1/gl"
)

// ShipData represents the matrix of the ship
type ShipData [][]int

// ShipType enum type
type ShipType int

const (
	// Ship1 type
	Ship1 ShipType = iota
	// Ship2 type
	Ship2
	// Ship3 type
	Ship3
	// Ship4 type
	Ship4
)

type ship struct {
	x, y  float32
	sType ShipType
	data  ShipData
}

// Ship interface
type Ship interface {
	Draw()
}

// NewShip creates a new Ship struct
func NewShip(m ShipData, t ShipType) Ship {
	return &ship{
		x:     0,
		y:     0,
		data:  m,
		sType: t,
	}
}

func (s *ship) Draw() {
	sHeight := float32(len(s.data))
	// sWidth := float32(len(s.data[0]))

	gl.PushMatrix()
	{
		gl.Color3f(0, 1, 0)
		for i := range s.data {
			y := float32(i)

			for j := range s.data[i] {
				x := float32(j)

				gl.Begin(gl.QUADS)
				{
					gl.Vertex2f(x, sHeight-y)
					gl.Vertex2f(x+1, sHeight-y)
					gl.Vertex2f(x+1, sHeight-y-1)
					gl.Vertex2f(x, sHeight-y-1)
				}
				gl.End()
			}
		}
	}
	gl.PopMatrix()
}
