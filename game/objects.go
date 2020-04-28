package game

import (
	"sync"

	"github.com/arielril/go-space-invaders/util"
)

var ships map[ShipType]ShipData
var colors [][]int
var car [][]int

var once sync.Once

// InitObjects initialize the objects for the game
func InitObjects() {
	ships = make(map[ShipType]ShipData, 0)

	once.Do(func() {
		ships[Ship1] = util.ParseFile("./templates/ship1.txt")
		ships[Ship2] = util.ParseFile("./templates/ship2.txt")
		ships[Ship3] = util.ParseFile("./templates/ship3.txt")
		ships[Ship4] = util.ParseFile("./templates/ship4.txt")

		colors = util.ParseFile("./templates/colors.txt")
		car = util.ParseFile("./templates/car.txt")
	})
}
