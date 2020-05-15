package game

import (
	"math/rand"

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

// measured in seconds
var shipSpeeds map[ShipType]float32 = map[ShipType]float32{
	Ship1: 12,
	Ship2: 9,
	Ship3: 5,
	Ship4: 15,
}

type ship struct {
	x, y  float32
	sType ShipType
	data  ObjectData
	scale float32
	speed float32
}

// Ship interface
type Ship interface {
	Object
	Die()
	GetX() float32
	GetY() float32
	Move()
	SetSpeed(fps float32) Ship
	RestartPos()
}

// NewShip creates a new Ship struct
func NewShip(m ObjectData, t ShipType) Ship {
	// (10 units * (1/fps)s ) in Xs
	speed := getShipSpeed(float32(maxFps), shipSpeeds[t])

	return &ship{
		x:     0,
		y:     0,
		data:  m,
		sType: t,
		scale: 1,
		speed: speed,
	}
}

func getShipSpeed(fps, shipSpeed float32) float32 {
	return float32(10*(1/fps)) / shipSpeed
}

func (s *ship) Draw() {
	sHeight := float32(len(s.data))

	gl.PushMatrix()
	{
		gl.Translatef(s.x, s.y, 0)
		gl.Scalef(s.scale, s.scale, 1)

		for i := range s.data {
			y := float32(i)

			for j, quadColor := range s.data[i] {
				ChangeColorFromInt(quadColor)

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

func (s *ship) Die() {
	var i int
	for idx, v := range gameShips {
		if v == s {
			i = idx
			break
		}
	}
	gameShips[i] = nil
}

func (s *ship) GetX() float32 {
	return s.x
}

func (s *ship) GetY() float32 {
	return s.y
}

func (s *ship) GetBoundingBox() BoundingBox {
	shipWidth := float64(len(s.data[0]))
	shipHeight := float64(len(s.data))

	bb := NewBoundingBox(
		float64(s.GetX()),
		float64(s.GetY()),
		shipWidth,
		shipHeight,
	)

	return bb
}

func (s *ship) SetSpeed(fps float32) Ship {
	s.speed = getShipSpeed(fps, shipSpeeds[s.sType])
	return s
}

func (s *ship) Move() {
	moveStep := s.speed
	newY := s.GetY() - moveStep
	if newY <= 0 {
		s.RestartPos()
	} else {
		s.SetY(newY)
	}
}

func (s *ship) RestartPos() {
	r := float32(rand.Intn(10))
	s.SetPos(r, 10)
}
