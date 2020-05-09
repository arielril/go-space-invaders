package game

type boundingBox struct {
	x, y, width, height float32
}

// BoundingBox interface
type BoundingBox interface {
	CheckColision(bbAgainst BoundingBox) bool
	GetPos() (x, y float32)
	GetSize() (width, height float32)
}

// NewBoundingBox creates a new bounding box pointer
func NewBoundingBox(x, y, width, height float32) BoundingBox {
	return &boundingBox{
		x:      x,
		y:      y,
		width:  width,
		height: height,
	}
}

// CheckColision
// verify if two bounding boxes has colidded
//
// Returns:
// 	- true if there is a colision
// 	- false if there is no colision
func (bb *boundingBox) CheckColision(bb2 BoundingBox) bool {
	bb2X, bb2Y := bb2.GetPos()
	bb2Width, bb2Height := bb2.GetSize()

	if bb.x < bb2X+bb2Width &&
		bb.x+bb.width > bb2X &&
		bb.y < bb2Y+bb2Height &&
		bb.y+bb.height > bb2Y {
		return true
	}

	return false
}

func (bb *boundingBox) GetPos() (x, y float32) {
	return bb.x, bb.y
}

func (bb *boundingBox) GetSize() (width, height float32) {
	return bb.width, bb.height
}
