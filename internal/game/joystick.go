package game

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	input "github.com/quasilyte/ebitengine-input"
	"github.com/quasilyte/gmath"
)

type VirtualJoystick struct {
	Center gmath.Vec
	Radius float64

	vec    gmath.Vec
	Active bool
}

func NewVirtualJoystick() *VirtualJoystick {
	return &VirtualJoystick{
		Center: gmath.Vec{X: 40, Y: 140}, // можешь потом подвинуть
		Radius: 32,
	}
}

func (v *VirtualJoystick) IsActive() bool {
	if v == nil {
		return false
	}

	return v.Active
}

func (v *VirtualJoystick) Update(_ *input.Handler) {
	if v == nil {
		return
	}

	v.Active = false
	v.vec = gmath.Vec{}

	ids := ebiten.AppendTouchIDs(nil)
	if len(ids) == 0 {
		return
	}

	x, y := ebiten.TouchPosition(ids[0])

	dx := float64(x) - v.Center.X
	dy := float64(y) - v.Center.Y

	dist := math.Hypot(dx, dy)
	if dist < 3 {
		return
	}

	if dist > v.Radius {
		s := v.Radius / dist
		dx *= s
		dy *= s
	}

	v.vec = gmath.Vec{
		X: dx / v.Radius,
		Y: dy / v.Radius,
	}

	v.Active = true
}

func (v *VirtualJoystick) Vector() gmath.Vec {
	return v.vec
}

func (v *VirtualJoystick) Draw(screen *ebiten.Image) {
	if v == nil {
		return
	}
	
	drawCircle(screen, v.Center.X, v.Center.Y, v.Radius, color.RGBA{50, 50, 80, 200})

	knobX := v.Center.X + v.vec.X*v.Radius
	knobY := v.Center.Y + v.vec.Y*v.Radius
	drawCircle(screen, knobX, knobY, 10, color.RGBA{200, 200, 255, 255})
}

func drawCircle(screen *ebiten.Image, cx, cy, r float64, col color.Color) {
	const steps = 32
	for i := 0; i < steps; i++ {
		a1 := float64(i) * 2 * math.Pi / steps
		a2 := float64(i+1) * 2 * math.Pi / steps
		x1 := cx + math.Cos(a1)*r
		y1 := cy + math.Sin(a1)*r
		x2 := cx + math.Cos(a2)*r
		y2 := cy + math.Sin(a2)*r
		ebitenutil.DrawLine(screen, x1, y1, x2, y2, col)
	}
}
