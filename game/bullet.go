package game

import (
	"github.com/go-gl/gl/v2.1/gl"
	glfw "github.com/go-gl/glfw/v3.3/glfw"
)

// Bullet interface
type Bullet interface {
	Object
	Move()
	SetY(y float32)
	GetY() float32
	GetX() float32
	Remove(w *glfw.Window) bool
	CheckHit(ship Ship) bool
}

type bullet struct {
	x, y  float32
	scale float32
}

// NewBullet returns a new bullet struct
func NewBullet() Bullet {
	return &bullet{
		x:     0,
		y:     0,
		scale: 1,
	}
}

// NewBulletWithPos returns a new bullet with the position set
func NewBulletWithPos(x, y float32) Bullet {
	b := NewBullet()

	return b.SetPos(x, y).(Bullet)
}

func (b *bullet) Draw() {
	gl.PushMatrix()
	{
		gl.Color3f(0, 0, 0)
		gl.Translatef(b.x, b.y, 0)
		gl.Scalef(b.scale, b.scale, 1)

		gl.Begin(gl.QUADS)
		{
			gl.Vertex2f(0, 0)
			gl.Vertex2f(1, 0)
			gl.Vertex2f(1, 1)
			gl.Vertex2f(0, 1)

			gl.Vertex2f(0, 1)
			gl.Vertex2f(1, 1)
			gl.Vertex2f(1, 2)
			gl.Vertex2f(0, 2)
		}
		gl.End()
	}
	gl.PopMatrix()
}

func (b *bullet) ResetScale() Object {
	b.scale = 1
	return b
}

func (b *bullet) SetScale(scale float32) Object {
	b.scale = scale
	return b
}

func (b *bullet) SetPos(x, y float32) Object {
	b.x = x
	b.y = y
	return b
}

func (b *bullet) SetY(y float32) {
	b.y = y
}

func (b *bullet) GetY() float32 {
	return b.y
}

func (b *bullet) GetX() float32 {
	return b.x
}

func (b *bullet) Move() {
	moveStep := 1.0 * float32(.15)
	b.SetY(b.GetY() + moveStep)
}

func removeBullet(b *bullet) {
	var idx int
	for i, v := range gameBullets {
		if b == v {
			idx = i
			break
		}
	}

	gameBullets[idx] = nil
}

func (b *bullet) Remove(w *glfw.Window) bool {
	if b.GetY() < -10 || b.GetY() > 10 {
		removeBullet(b)
		return true
	}

	if b.GetX() > 10 {
		removeBullet(b)
		return true
	}

	return false
}

func (b *bullet) CheckHit(ship Ship) bool {
	// bX := b.GetX()
	// bY := b.GetY()

	// sX := ship.GetX()
	// sY := ship.GetY()

	return false
}

func (b *bullet) GetBoundingBox() BoundingBox {
	bb := NewBoundingBox(0, 0, 0, 0)

	return bb
}
