package game

import "github.com/go-gl/gl/v2.1/gl"

// MoveDirection direction for the car to move
type MoveDirection int

const (
	// leftMove move to the left
	leftMove = iota
	// rightMove move to the right
	rightMove
)

// Car interface
type Car interface {
	Object
	SetX(x float32) Car
	GetX() float32
	GetY() float32
	MoveRight() Car
	MoveLeft() Car
	Shoot()
	ResetPos() Car
}

type car struct {
	x, y  float32
	data  ObjectData
	scale float32
}

// NewCar creates a new car
func NewCar(m ObjectData) Car {
	return &car{
		x:     0,
		y:     0,
		data:  m,
		scale: 1,
	}
}

func (c *car) Draw() {
	carHeight := float32(len(c.data))

	gl.PushMatrix()
	{
		gl.Translatef(c.x, c.y, 0)
		gl.Scalef(c.scale, c.scale, 1)

		for i := range c.data {
			y := float32(i)

			for j, pixColor := range c.data[i] {
				ChangeColorFromInt(pixColor)

				x := float32(j)

				gl.Begin(gl.QUADS)
				{
					gl.Vertex2f(x, carHeight-y)
					gl.Vertex2f(x+1, carHeight-y)
					gl.Vertex2f(x+1, carHeight-y-1)
					gl.Vertex2f(x, carHeight-y-1)
				}
				gl.End()
			}
		}
	}
	gl.PopMatrix()
}

func (c *car) SetPos(x, y float32) Object {
	c.x = x
	c.y = y

	return c
}

func (c *car) SetScale(scale float32) Object {
	c.scale = scale
	return c
}

func (c *car) ResetScale() Object {
	c.scale = 1
	return c
}

func (c *car) SetX(x float32) Car {
	c.x = x
	return c
}

func (c *car) GetX() float32 {
	return c.x
}

func (c *car) GetY() float32 {
	return c.y
}

func (c *car) move(dir MoveDirection) Car {
	moveStep := 1.0 * float32(.3)
	var newPos float32

	switch dir {
	case leftMove:
		newPos = c.GetX() - moveStep
		break
	case rightMove:
		newPos = c.GetX() + moveStep
		break
	}

	if newPos < 0 {
		newPos = 10
	}

	if newPos > 10 {
		newPos = 0
	}

	c.SetX(newPos)

	return c
}

func (c *car) MoveLeft() Car {
	return c.move(leftMove)
}

func (c *car) MoveRight() Car {
	return c.move(rightMove)
}

func (c *car) Shoot() {
	bullet := NewBulletWithPos(
		c.GetX()+.35,
		c.GetY(),
	)

	AddShoot(bullet)
}

func (c *car) GetBoundingBox() BoundingBox {
	carWidth := float64(len(c.data[0]))
	carHeight := float64(len(c.data))

	bb := NewBoundingBox(
		float64(c.GetX()),
		float64(c.GetY()),
		carWidth,
		carHeight,
	)

	return bb
}

func (c *car) ResetPos() Car {
	c.SetPos(5, 0)
	return c
}
