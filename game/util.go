package game

import glfw "github.com/go-gl/glfw/v3.3/glfw"

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
