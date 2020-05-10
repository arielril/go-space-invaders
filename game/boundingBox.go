package game

import (
	"fmt"
	"math"
)

type boundingBox struct {
	xCenter, yCenter, xWidth, yWidth float32
}

// BoundingBox interface
type BoundingBox interface {
	CheckCollision(bbAgainst BoundingBox) bool
	GetCenterPoint() (x, y float32)
	GetHalfWidthPoint() (x, y float32)
}

// NewBoundingBox creates a new bounding box pointer
func NewBoundingBox(x, y, width, height float32) BoundingBox {
	xCenter := (x + (width / 2))  // * gameScale
	yCenter := (y + (height / 2)) // * gameScale

	xWidth := (width / 2)  // * gameScale
	yWidth := (height / 2) // * gameScale

	return &boundingBox{
		xCenter, yCenter, xWidth, yWidth,
	}
}

// CheckCollision
// verify if two bounding boxes has colidded
//
// Returns:
// 	- true if there is a collision
// 	- false if there is no collision
func (bb *boundingBox) CheckCollision(bb2 BoundingBox) bool {
	bbX, bbY := bb.GetCenterPoint()
	bbXWidth, bbYWidth := bb.GetHalfWidthPoint()

	bb2X, bb2Y := bb2.GetCenterPoint()
	bb2XWidth, bb2YWidth := bb2.GetHalfWidthPoint()

	fmt.Printf("BBc => (%v, %v)\n", bbX, bbY)
	fmt.Printf("BB2c => (%v, %v)\n", bb2X, bb2Y)

	fmt.Printf("BBH => (%v, %v)\n", bbXWidth, bbYWidth)
	fmt.Printf("BB2H => (%v, %v)\n\n", bb2XWidth, bb2YWidth)

	if math.Abs(float64(bbX-bb2X)) > float64(bbXWidth+bb2XWidth) {
		return false
	}

	if math.Abs(float64(bbY-bb2Y)) > float64(bbYWidth+bb2YWidth) {
		return false
	}

	return true
}

func (bb *boundingBox) GetCenterPoint() (x, y float32) {
	x = bb.xCenter
	y = bb.yCenter
	return x, y
}

func (bb *boundingBox) GetHalfWidthPoint() (x, y float32) {
	x = bb.xWidth
	y = bb.yWidth

	return x, y
}
