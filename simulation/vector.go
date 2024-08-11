package simulation

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	degrees120 = 2.0944 // 120 degrees in radians
)

type Vector struct {
	X, Y float64
}

func NewVector(x, y float64) Vector {
	return Vector{X: x, Y: y}
}

func (v Vector) Add(other Vector) Vector {
	return Vector{X: v.X + other.X, Y: v.Y + other.Y}
}

func (v Vector) Rotate(angle float64) Vector {
	sin, cos := math.Sincos(angle)
	return Vector{
		X: v.X*cos - v.Y*sin,
		Y: v.X*sin + v.Y*cos,
	}
}

func DrawTriangle(screen *ebiten.Image, center Vector, size float64, angle float64, c color.Color) {
	v1 := NewVector(size, 0).Rotate(angle)
	v2 := NewVector(size, 0).Rotate(angle + degrees120)
	v3 := NewVector(size, 0).Rotate(angle - degrees120)

	p1 := center.Add(v1)
	p2 := center.Add(v2)
	p3 := center.Add(v3)

	vector.StrokeLine(screen, float32(p1.X), float32(p1.Y), float32(p2.X), float32(p2.Y), 1, c, false)
	vector.StrokeLine(screen, float32(p2.X), float32(p2.Y), float32(p3.X), float32(p3.Y), 1, c, false)
	vector.StrokeLine(screen, float32(p3.X), float32(p3.Y), float32(p1.X), float32(p1.Y), 1, c, false)
}
