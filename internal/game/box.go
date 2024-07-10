package game

import (
	input "github.com/quasilyte/ebitengine-input"
	"github.com/quasilyte/gmath"
)

type Box struct {
	Speed       float64
	Object      *Object
	Collider    *Collider
	LastPos     gmath.Vec
	HasCollided bool
}

func NewBox(speed float64, object *Object,
	collider *Collider, inputHandler *input.Handler) *Box {
	lastPos := object.Pos

	return &Box{
		Speed:    speed,
		Object:   object,
		Collider: collider,
		LastPos:  lastPos,
	}
}

func (b *Box) Update() {
	//b.Collider.Update()
}
