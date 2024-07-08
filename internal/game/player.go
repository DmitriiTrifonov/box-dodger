package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	input "github.com/quasilyte/ebitengine-input"
	"github.com/quasilyte/gmath"
	"log"
)

type Player struct {
	Speed       float64
	Object      *Object
	Collider    *Collider
	Input       *input.Handler
	LastPos     *gmath.Vec
	HasCollided bool
}

func NewPlayer(speed float64, object *Object,
	collider *Collider, inputHandler *input.Handler) *Player {
	lastPos := object.Pos

	return &Player{
		Speed:    speed,
		Object:   object,
		Collider: collider,
		Input:    inputHandler,
		LastPos:  lastPos,
	}
}

func (p *Player) Update(actualTPS float64) {
	p.UpdateLastPos()
	p.Move(actualTPS)
	p.UpdateCollider()
	log.Println("cur", p.Object.Pos, "last", p.LastPos)
}

func (p *Player) SetLastPos() {
	p.Object.SetPos(p.LastPos)
	p.UpdateCollider()
}

func (p *Player) UpdateCollider() {
	p.Collider.Update(p.Object.Pos)
}

func (p *Player) UpdateLastPos() {
	if !p.HasCollided {
		log.Println("change last")
		p.LastPos = p.Object.Pos
	}
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

	if !p.HasCollided {
		p.Object.Pos.X += vec.X
		p.Object.Pos.Y += vec.Y
	} else {
		p.Object.Pos.X = p.LastPos.X
		p.Object.Pos.Y = p.LastPos.Y
		p.HasCollided = false
	}

}

func (p *Player) Draw(screen *ebiten.Image) {
	p.Object.Draw(screen)
}
