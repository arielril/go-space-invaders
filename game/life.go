package game

import (
	"github.com/go-gl/gl/v2.1/gl"
)

// Life of the player
type Life interface {
	Draw()
	SetScale(scale float32) Life
}

type life struct {
	x, y  float32
	scale float32
	data  ObjectData
}

var playerLives []Life

// NewLife create a new life object
func NewLife(m ObjectData, x, y float32) Life {
	return &life{
		x:     x,
		y:     y,
		scale: 1,
		data:  m,
	}
}

func (l *life) Draw() {
	lifeHeight := float32(len(l.data))

	gl.PushMatrix()
	{
		gl.Translatef(l.x, l.y, 0)
		gl.Scalef(l.scale, l.scale, 1)

		for i := range l.data {
			y := float32(i)

			for j, quadColor := range l.data[i] {
				x := float32(j)
				ChangeColorFromInt(quadColor)

				gl.Begin(gl.QUADS)
				{
					gl.Vertex2f(x, lifeHeight-y)
					gl.Vertex2f(x+1, lifeHeight-y)
					gl.Vertex2f(x+1, lifeHeight-y-1)
					gl.Vertex2f(x, lifeHeight-y-1)
				}
				gl.End()
			}
		}
	}
	gl.PopMatrix()
}

func (l *life) SetScale(scale float32) Life {
	l.scale = scale
	return l
}

// KillPlayer removes one life of the player
func KillPlayer() {
	playerLives = playerLives[:len(playerLives)-1]
}

// IsAlive verifies if the player is alive
func IsAlive() bool {
	return len(playerLives) > 0
}
