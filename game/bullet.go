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
	Hit()
}

type bullet struct {
	x, y  float32
	scale float32
	speed float32
	data  ObjectData
}

// NewBullet returns a new bullet struct
func NewBullet(d ObjectData) Bullet {
	speed := (10 * (1 / float32(maxFps))) / 3

	return &bullet{
		x:     0,
		y:     0,
		scale: 1,
		speed: speed,
		data:  d,
	}
}

// NewBulletWithPos returns a new bullet with the position set
func NewBulletWithPos(x, y float32) Bullet {
	b := NewBullet(bulletData)

	return b.SetPos(x, y).(Bullet)
}

func (b *bullet) Draw() {
	bulletHeight := float32(len(b.data))

	gl.PushMatrix()
	{
		gl.Translatef(b.x, b.y, 0)
		gl.Scalef(b.scale, b.scale, 1)

		for i := range b.data {
			y := float32(i)

			for j, quadColor := range b.data[i] {
				x := float32(j)
				ChangeColorFromInt(quadColor)

				gl.Begin(gl.QUADS)
				{
					gl.Vertex2f(x, bulletHeight-y)
					gl.Vertex2f(x+1, bulletHeight-y)
					gl.Vertex2f(x+1, bulletHeight-y-1)
					gl.Vertex2f(x, bulletHeight-y-1)
				}
				gl.End()
			}
		}
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
	moveStep := b.speed
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
	if b.GetX() < 0 || b.GetX() > 20 {
		removeBullet(b)
		return true
	}

	if b.GetY() > 10 {
		removeBullet(b)
		return true
	}

	return false
}

func (b *bullet) Hit() {
	removeBullet(b)
}

func (b *bullet) GetBoundingBox() BoundingBox {
	bWidth := float64(len(b.data[0]))
	bHeight := float64(len(b.data))

	bb := NewBoundingBox(
		float64(b.GetX()),
		float64(b.GetY()),
		bWidth,
		bHeight,
	)

	return bb
}
