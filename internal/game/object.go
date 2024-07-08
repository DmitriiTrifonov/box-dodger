package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gmath"
)

type Object struct {
	Sprite       *Sprite
	Pos          gmath.Vec
	IsStatic     bool
	StartAnimIdx int
}

func (o *Object) SetPos(newPos gmath.Vec) {
	o.Pos = newPos
}

func (o *Object) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(o.Pos.X, o.Pos.Y)
	o.Sprite.Draw(screen, opts)
}

func (o *Object) GetSprite() *Sprite {
	return o.Sprite
}
