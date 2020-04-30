package game

import (
	"github.com/go-gl/gl/v2.1/gl"
)

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
	data  ObjectData
	scale float32
}

// Ship interface
type Ship interface {
	Object
	Die() Ship
}

// NewShip creates a new Ship struct
func NewShip(m ObjectData, t ShipType) Ship {
	return &ship{
		x:     0,
		y:     0,
		data:  m,
		sType: t,
		scale: 1,
	}
}

func (s *ship) Draw() {
	sHeight := float32(len(s.data))

	gl.PushMatrix()
	{
		gl.Translatef(s.x, s.y, 0)
		gl.Scalef(s.scale, s.scale, 1)

		for i := range s.data {
			y := float32(i)

			for j, pixColor := range s.data[i] {
				ChangeColorFromInt(pixColor)

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

func (s *ship) SetPos(x, y float32) Object {
	s.x = x
	s.y = y

	return s
}

func (s *ship) SetX(x float32) Ship {
	s.x = x
	return s
}

func (s *ship) SetY(y float32) Ship {
	s.y = y
	return s
}

func (s *ship) SetScale(sc float32) Object {
	s.scale = sc
	return s
}

func (s *ship) ResetScale() Object {
	s.scale = 1
	return s
}

func (s *ship) Die() Ship {
	var i int
	for idx, v := range gameShips {
		if v == s {
			i = idx
			break
		}
	}
	gameShips = gameShips[: i-1 : i+1]
	return s
}
