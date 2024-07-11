package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gmath"
	"time"
)

type Box struct {
	Speed       float64
	Object      *Object
	Collider    *Collider
	LastPos     gmath.Vec
	HasCollided bool
	Target      gmath.Vec
	rand        *gmath.Rand
}

func NewBox(speed float64, object *Object,
	collider *Collider) *Box {
	lastPos := object.Pos

	rand := &gmath.Rand{}
	rand.SetSeed(time.Now().UnixNano())
	direction := rand.Offset(24.0, 144.0)

	return &Box{
		Speed:    speed,
		Object:   object,
		Collider: collider,
		LastPos:  lastPos,
		Target:   direction,
		rand:     rand,
	}
}

func (b *Box) Update(tps int, g *Game) {
	b.Speed += 1
	b.Object.Pos = b.Object.Pos.MoveTowards(b.Target, b.Speed/float64(tps))
	b.Collider.Update(b.Object.Pos)
	if b.Collider.CheckCollisionsWithTag(collisionTagPlayer) {
		if time.Since(g.StartTime) > 2*time.Second {
			g.IsGameOver = true
			g.TotalTime = time.Since(g.StartTime)
		}
	}

	if b.Collider.CheckCollisionsWithTag(collisionTagWalls) {
		b.Target = b.setNewTarget()
	}

	if b.Object.Pos.EqualApprox(b.Target) {
		b.Target = b.setNewTarget()
	}
}

func (b *Box) Draw(screen *ebiten.Image) {
	b.Object.Draw(screen)
}

func (b *Box) setNewTarget() gmath.Vec {
	x := b.rand.FloatRange(24, 360)
	y := b.rand.FloatRange(24, 180)

	return gmath.Vec{x, y}
}
