package game

import (
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

func getBulletFromShipHit(shipBoundingBox BoundingBox) Bullet {
	for _, bullet := range gameBullets {
		if bullet == nil {
			continue
		}

		bulletBoundingBox := bullet.GetBoundingBox()
		if bulletBoundingBox.CollidedWith(shipBoundingBox) {
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
		shipBoundingBox := ship.GetBoundingBox()

		if len(gameBullets) > 0 {
			// * bullet hitting some ship
			bulletThatHitShip := getBulletFromShipHit(shipBoundingBox)
			if bulletThatHitShip != nil {
				ship.Die()              // * bye bye ship :'(
				bulletThatHitShip.Hit() // * nice work bullet! Go home now :D
				continue
			}
		}

		// * ship hitting the car
		carBoundingBox := GetCar().GetBoundingBox()
		shipHitTheCar := shipBoundingBox.CollidedWith(carBoundingBox)
		if shipHitTheCar {
			KillPlayer()
			GetCar().ResetPos()
			ship.RestartPos()
		}
	}
}

// PlayerHasWon returns wheather the player won the game or not
func PlayerHasWon() bool {
	if len(gameShips) <= 0 {
		return true
	}
	return false
}
