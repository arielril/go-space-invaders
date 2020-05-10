package game

type boundingBox struct {
	topLeftX, topLeftY, bottomRightX, bottomRightY float32
}

// BoundingBox interface
type BoundingBox interface {
	CheckCollision(bbAgainst BoundingBox) bool
	GetValues() (tlX, tlY, brX, brY float32)
}

// NewBoundingBox creates a new bounding box pointer
func NewBoundingBox(x, y, width, height float32) BoundingBox {
	topLeftX := x
	topLeftY := (y + height)

	bottomRightX := (x + width)
	bottomRightY := y

	return &boundingBox{
		topLeftX, topLeftY, bottomRightX, bottomRightY,
	}
}

func (bb *boundingBox) GetValues() (tlX, tlY, brX, brY float32) {
	return bb.topLeftX, bb.topLeftY, bb.bottomRightX, bb.bottomRightY
}

// CheckCollision
// verify if two bounding boxes has colidded
//
// Returns:
// 	- true if there is a collision
// 	- false if there is no collision
func (bb *boundingBox) CheckCollision(bb2 BoundingBox) bool {
	// TODO: check this function
	baseTlX, baseTlY, _, _ := bb.GetValues()
	againstTlX, _, againstBrX, againstBrY := bb2.GetValues()
	/*
		collisionX := baseTlX+baseBrX >= againstTlX && againstTlX+againstBrX >= baseBrX
		collisionY := baseTlY+baseBrY >= againstTlY && againstTlY+againstBrY >= baseBrY

		fmt.Printf("coll X %v\n", collisionX)
		fmt.Printf("coll Y %v\n", collisionY)

		return collisionX && collisionY
	*/

	return (baseTlX >= againstTlX && baseTlX <= againstBrX) && baseTlY >= againstBrY
}
