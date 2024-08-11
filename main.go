package main

import (
	"flag"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafayettegabe/abm-flocking-simulation/constants"
	"github.com/lafayettegabe/abm-flocking-simulation/simulation"
)

func main() {
	numBirds := flag.Int("birds", 3, "number of birds")
	speed := flag.Float64("speed", 3.0, "speed of birds")
	flag.Parse()

	sim := simulation.NewSimulation(*numBirds, *speed)

	ebiten.SetWindowSize(constants.ScreenWidth, constants.ScreenHeight)
	ebiten.SetWindowTitle("Flocking Simulation")
	if err := ebiten.RunGame(sim); err != nil {
		panic(err)
	}
}
