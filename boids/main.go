package main

import (
	"image/color"
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenHeight, screenWidth = 360, 640
	boidCount = 250
	viewRadius = 20
	adjRate = 0.15 // adjustment rate
)

var (
	green = color.RGBA{10, 255, 50, 255}
	boids [boidCount]*Boid
	boidIdMatrix [screenWidth+1][screenHeight+1]int
	lock = sync.RWMutex{}
)

type Game struct{}

//update the screen
func (g *Game) Update() error {
	return nil
}

//draw a square with 4 coordinates using the boids position on the screen
func (g *Game) Draw(screen *ebiten.Image) {

	for _, boid := range boids {
		screen.Set(int(boid.position.x+1), int(boid.position.y+1), green) // one point of the diagonal
		screen.Set(int(boid.position.x-1), int(boid.position.y-1), green) // another point of the diagonal
		screen.Set(int(boid.position.x+1), int(boid.position.y-1), green)
		screen.Set(int(boid.position.x-1), int(boid.position.y+1), green)
	}
}

// create a layout of the window
func (g *Game) Layout(_, _ int) (w, h int) {
	return screenWidth, screenHeight
}

func main() {
	// create a boid matrix
	for i, row := range boidIdMatrix {
		for j := range row {
			boidIdMatrix[i][j] = -1
		}
	}

	// create each boid and spawn it
	for i:= 0; i < boidCount; i++ {
		createBoid(i)
	}

	// set the ebiten window or graphics window size
	ebiten.SetWindowSize(screenWidth, screenHeight)
	// set the winows title
	ebiten.SetWindowTitle("Boids Simulation")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

}