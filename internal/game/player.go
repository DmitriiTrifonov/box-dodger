package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	input "github.com/quasilyte/ebitengine-input"
	"github.com/quasilyte/gmath"
)

type Player struct {
	Speed       float64
	Object      *Object
	Collider    *Collider
	Input       *input.Handler
	LastPos     gmath.Vec
	HasCollided bool
}

func NewPlayer(speed float64, object *Object,
	collider *Collider, inputHandler *input.Handler) *Player {
	lastPos := object.Pos

	AddCollisionToTag(collisionTagPlayer, collider)

	return &Player{
		Speed:    speed,
		Object:   object,
		Collider: collider,
		Input:    inputHandler,
		LastPos:  lastPos,
	}
}

func (p *Player) Update(tps int) {
	p.UpdateLastPos()
	p.Move(tps)
}

func (p *Player) SetLastPos() {
	p.Object.SetPos(p.LastPos)
}

func (p *Player) UpdateCollider(vec gmath.Vec) {
	p.Collider.Update(vec)
}

func (p *Player) UpdateLastPos() {
	if !p.HasCollided {
		p.LastPos = p.Object.Pos
	}
}

func (p *Player) MoveOld(actualTPS int) {
	//velocity := gmath.Vec{}

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

	vec = vec.ClampLen(1).Mulf(p.Speed / float64(actualTPS))

	if !p.HasCollided {
		p.Object.Pos = p.Object.Pos.Add(vec)
	} else {
		p.Object.Pos = p.LastPos
	}
}

func (p *Player) Move(tps int) {
	// 1) Check input
	// 2) Update vector
	// 3) Update collider
	// 4) Check for collisions
	// 5) Update movement

	vec := p.getInputAxis().ClampLen(1).Mulf(p.Speed / float64(tps))

	position := p.Object.Pos.Add(vec)

	p.UpdateCollider(position)

	if p.Collider.CheckCollisionsWithTag(collisionTagWalls) {
		p.HasCollided = true
		position = p.LastPos

		return
	}

	p.HasCollided = false
	p.Object.Pos = position
}

func (p *Player) getInputAxis() gmath.Vec {
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

	return gmath.Vec{X: x, Y: y}
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.Object.Draw(screen)
}
