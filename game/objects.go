package game

import (
	"sync"

	"github.com/go-gl/gl/v2.1/gl"

	"github.com/arielril/go-space-invaders/util"
)

// ObjectData represents the matrix of the object
type ObjectData [][]int

var ships map[ShipType]ObjectData
var colors [][]int
var carData [][]int
var lifeData [][]int

var once sync.Once

// Object interface that represents the game objects
type Object interface {
	Draw()
	SetPos(x, y float32) Object
	SetScale(scale float32) Object
	ResetScale() Object
	GetBoundingBox() BoundingBox
}

// InitObjects initialize the objects for the game
func InitObjects() {
	ships = make(map[ShipType]ObjectData, 0)

	once.Do(func() {
		ships[Ship1] = util.ParseFile("./templates/ship1.txt")
		ships[Ship2] = util.ParseFile("./templates/ship2.txt")
		ships[Ship3] = util.ParseFile("./templates/ship3.txt")
		ships[Ship4] = util.ParseFile("./templates/ship4.txt")

		colors = util.ParseFile("./templates/colors.txt")
		carData = util.ParseFile("./templates/car.txt")
		lifeData = util.ParseFile("./templates/life.txt")
	})
}

// ChangeColorFromInt receives an int and translate the value to a RGB color
func ChangeColorFromInt(c int) {
	for _, v := range colors {
		if v[0] == c {
			r := float32(v[1])
			g := float32(v[2])
			b := float32(v[3])

			gl.Color3f(r, g, b)
		}
	}
}
