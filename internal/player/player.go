package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	input "github.com/quasilyte/ebitengine-input"
	"github.com/quasilyte/gmath"
	"github.com/solarlune/goaseprite"
	"log"
)

const (
	ActionMoveLeft input.Action = iota
	ActionMoveRight
	ActionMoveUp
	ActionMoveDown
)

type Player struct {
	Speed  float64
	Sprite *goaseprite.File
	Anim   *goaseprite.Player
	Img    *ebiten.Image
	Pos    *gmath.Vec
	Input  *input.Handler
}

func (p *Player) Update(actualTPS float64) {
	p.Move(actualTPS)
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
	log.Println("speed", vec.Len())

	p.Pos.X += vec.X
	p.Pos.Y += vec.Y
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.Pos.X, p.Pos.Y)
	screen.DrawImage(p.Img, op)
}
