package game

import (
	"fmt"

	"github.com/go-gl/gl/v2.1/gl"
)

type boundingBox struct {
	centerX, centerY      float64
	halfWidth, halfHeight float64
	scale                 float32
}

// BoundingBox interface
type BoundingBox interface {
	CollidedWith(bbAgainst BoundingBox) bool
	GetValues() (centerX, centerY, halfWidth, halfHeight float64)
	Draw()
}

// NewBoundingBox creates a new bounding box pointer
func NewBoundingBox(x, y, width, height float64) BoundingBox {
	halfWidth := width / 2
	halfHeight := height / 2

	centerX := x + halfWidth
	centerY := y + halfHeight

	return &boundingBox{
		centerX:    centerX,
		centerY:    centerY,
		halfWidth:  halfWidth,
		halfHeight: halfHeight,
		scale:      gameScale,
	}
}

func (bb *boundingBox) GetValues() (centerX, centerY, halfWidth, halfHeight float64) {
	return bb.centerX, bb.centerY, bb.halfWidth, bb.halfHeight
}

/*
func (a *AABB) Intersect(b *AABB) bool {
	if b.center.x+b.half.x < a.center.x-a.half.x {
			return false
	}
	if b.center.y+b.half.y < a.center.y-a.half.y {
			return false
	}
	if b.center.x-b.half.x > a.center.x+a.half.x {
			return false
	}
	if b.center.y-b.half.y > a.center.y+a.half.y {
			return false
	}

	return true
}
*/

// CollidedWith returns if an object collided with another object
func (bb *boundingBox) CollidedWith(bb2 BoundingBox) bool {
	// values of the base object
	centerX, centerY, halfWidth, halfHeight := bb.GetValues()

	// values from the against object
	aCenterX, aCenterY, aHalfWidth, aHalfHeight := bb2.GetValues()

	if aCenterX+aHalfWidth < centerX-halfWidth {
		return false
	}
	if aCenterY+aHalfHeight < centerY-halfHeight {
		return false
	}

	if centerX+halfWidth < aCenterX-aHalfWidth {
		return false
	}
	if centerY+halfHeight < aCenterY-aHalfHeight {
		return false
	}
	fmt.Printf(
		"Ac = (%v, %v) | Hw = %v | Hh = %v\n",
		centerX,
		centerY,
		halfWidth,
		halfHeight,
	)
	fmt.Printf("Bc = (%v, %v) | Hw = %v | Hh = %v\n", aCenterX, aCenterY, aHalfWidth, aHalfHeight)

	return true

	/* OLD

	xAbs := math.Abs(centerX - aCenterX)
	wSum := halfWidth + aHalfWidth

	if xAbs > wSum {
		return false
	}

	yAbs := math.Abs(centerY - aCenterY)
	hSum := halfHeight + aHalfHeight

	if yAbs > hSum {
		return false
	}

	fmt.Printf("FST check: %v > %v\n", xAbs, wSum)
	fmt.Printf("SND check: %v > %v\n", yAbs, hSum)


	return true
	*/
}

func (bb *boundingBox) Draw() {
	zeroX := float32(bb.centerX - bb.halfWidth)
	zeroY := float32(bb.centerY - bb.halfHeight)

	gl.PushMatrix()
	{
		gl.Color3f(0, 1, 0)

		gl.Translatef(zeroX, zeroY, 0)
		gl.Scalef(bb.scale, bb.scale, 1)

		gl.Begin(gl.QUADS)
		{
			gl.Vertex2f(-1, -1)
			gl.Vertex2f(-1, (float32(bb.halfHeight)*2)+1)
			gl.Vertex2f((float32(bb.halfWidth)*2)+1, (float32(bb.halfHeight)*2)+1)
			gl.Vertex2f((float32(bb.halfWidth)*2)+1, -1)
		}
		gl.End()
	}
	gl.PopMatrix()
}
