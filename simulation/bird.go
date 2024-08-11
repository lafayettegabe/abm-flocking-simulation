package simulation

import (
	"image/color"
	"math"
	"math/rand/v2"
	"sync"

	"github.com/lafayettegabe/abm-flocking-simulation/constants"
)

const (
	BirdSize = 10
)

type Bird struct {
	X, Y  float64
	Angle float64
	Color color.Color
}

func (b *Bird) Update(speed float64, wg *sync.WaitGroup) {
	defer wg.Done()

	b.X += math.Cos(b.Angle) * speed
	b.Y += math.Sin(b.Angle) * speed

	b.X = math.Mod(b.X+constants.ScreenWidth, constants.ScreenWidth)
	b.Y = math.Mod(b.Y+constants.ScreenHeight, constants.ScreenHeight)
}

func NewBird() *Bird {

	return &Bird{
		X:     rand.Float64() * constants.ScreenWidth,
		Y:     rand.Float64() * constants.ScreenHeight,
		Angle: rand.Float64() * 2 * math.Pi,
		Color: color.NRGBA{
			R: uint8(rand.Float64() * 255),
			G: uint8(rand.Float64() * 255),
			B: uint8(rand.Float64() * 255),
			A: 255,
		},
	}
}
