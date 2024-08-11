package simulation

import (
	"fmt"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lafayettegabe/abm-flocking-simulation/constants"
)

type Simulation struct {
	Birds []*Bird
	Speed float64
}

func (s *Simulation) Update() error {
	var wg sync.WaitGroup
	wg.Add(len(s.Birds))

	for _, bird := range s.Birds {
		go bird.Update(s.Speed, &wg)
	}

	wg.Wait()
	return nil
}

func (s *Simulation) Draw(screen *ebiten.Image) {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	chunkSize := 10
	for i := 0; i < len(s.Birds); i += chunkSize {
		end := i + chunkSize
		if end > len(s.Birds) {
			end = len(s.Birds)
		}

		wg.Add(1)
		go func(birds []*Bird) {
			defer wg.Done()
			for _, bird := range birds {
				center := NewVector(bird.X, bird.Y)

				mutex.Lock()
				DrawTriangle(screen, center, BirdSize, bird.Angle, bird.Color)
				mutex.Unlock()
			}
		}(s.Birds[i:end])
	}

	wg.Wait()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Birds: %d, Speed: %.2f", len(s.Birds), s.Speed))
}

func (s *Simulation) Layout(outsideWidth, outsideHeight int) (int, int) {
	return constants.ScreenWidth, constants.ScreenHeight
}

func NewSimulation(numBirds int, speed float64) *Simulation {
	birds := make([]*Bird, numBirds)
	for i := range birds {
		birds[i] = NewBird()
	}

	return &Simulation{
		Birds: birds,
		Speed: speed,
	}
}
