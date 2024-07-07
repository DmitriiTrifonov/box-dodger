package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	input "github.com/quasilyte/ebitengine-input"
	"github.com/quasilyte/gmath"
)

type Player struct {
	Speed    float64
	Object   *Object
	Collider *Collider
	Input    *input.Handler
}

func (p *Player) Update(actualTPS float64) {
	p.Move(actualTPS)
	p.UpdateCollider()
}

func (p *Player) UpdateCollider() {
	p.Collider.Update(p.Object.Pos)
}

func (p *Player) Move(actualTPS float64) {
	if actualTPS == 0 {
		actualTPS = float64(ebiten.TPS())
	}

	x, y := 0.0, 0.0

	if p.Input.ActionIsPressed(ActionMoveLeft) {
		x = -1
	}

	if p.Input.ActionIsPressed(ActionMoveRight) {
		x = 1
	}

	if p.Input.ActionIsPressed(ActionMoveUp) {
		y = -1
	}

	if p.Input.ActionIsPressed(ActionMoveDown) {
		y = 1
	}

	vec := gmath.Vec{
		X: x,
		Y: y,
	}

	vec = vec.ClampLen(1).Mulf(p.Speed / actualTPS)

	p.Object.Pos.X += vec.X
	p.Object.Pos.Y += vec.Y
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.Object.Draw(screen)
}
