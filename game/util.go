package game

import (
	"fmt"

	glfw "github.com/go-gl/glfw/v3.3/glfw"
)

// optimizeGame
//
// remove nil objects and clean the structs of the game
func optimizeGame() {
	// remove nil bullets
	auxGameBullets := make([]Bullet, 0)
	for _, v := range gameBullets {
		if v != nil {
			auxGameBullets = append(auxGameBullets, v)
		}
	}
	gameBullets = auxGameBullets

	// remove nil ships
	auxGameShips := make([]Ship, 0)
	for _, v := range gameShips {
		if v != nil {
			auxGameShips = append(auxGameShips, v)
		}
	}
	gameShips = auxGameShips
}

// removeBulletsFromGame
//
// remove the invalid bullets from the
// gameBullets slice to open new space for new bullets
func removeBulletsFromGame(w *glfw.Window) {
	// check if the bullet is gone and remove it
	for _, v := range gameBullets {
		if v == nil {
			continue
		}
		v.Remove(w)
	}
}

// GetCar returns the car object
func GetCar() Car {
	return gameCar
}

// AddShoot add a new bullet to the game
func AddShoot(b Bullet) {
	if len(gameBullets) < maxShoots {
		gameBullets = append(gameBullets, b)
	}
}

func getBulletFromShipHit(s Ship) Bullet {
	shipBoundingBox := s.GetBoundingBox()

	for _, bullet := range gameBullets {
		if bullet == nil {
			continue
		}

		bulletBoundingBox := bullet.GetBoundingBox()
		if bulletBoundingBox.CollidedWith(shipBoundingBox) {
			// fmt.Printf("Apos = (%v, %v)\n", bullet.GetX(), bullet.GetY())
			// fmt.Printf("Bpos = (%v, %v)\n", s.GetX(), s.GetY())
			return bullet
		}
	}

	return nil
}

func doCollisions() {
	/*
		* Check for bullet -> ship collisions
		* Check for ship -> car collisions

		The mainstream is to check for ship collisions!
	*/

	for _, ship := range gameShips {
		if len(gameBullets) > 0 {
			// * bullet hitting some ship
			bulletThatHitShip := getBulletFromShipHit(ship)
			if bulletThatHitShip != nil {
				fmt.Printf("\nship hit %#v\n", ship)
				fmt.Printf("bullet %#v\n", bulletThatHitShip)
				fmt.Println("ship hit with bullet")

				ship.Die()              // * bye bye ship :'(
				bulletThatHitShip.Hit() // * nice work bullet! Go home now :D

				continue
			}
		}

		// * ship hitting the car
		shipHitTheCar := false
		if shipHitTheCar {
			fmt.Println("ship hit the car")
			KillPlayer()
			GetCar().ResetPos()
		}
	}
}
