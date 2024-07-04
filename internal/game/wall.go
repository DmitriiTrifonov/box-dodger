package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gmath"
	"github.com/solarlune/goaseprite"
)

type Wall struct {
	Sprite   *goaseprite.File
	Anim     *goaseprite.Player
	Img      *ebiten.Image
	Pos      *gmath.Vec
	Disabled bool
}

func (w *Wall) Update() {
	// Add collision check
}

func (w *Wall) Draw(screen *ebiten.Image) {
	if !w.Disabled {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(w.Pos.X, w.Pos.Y)
		screen.DrawImage(w.Img, op)
	}

}

func (w *Wall) RenderBackground(screen *ebiten.Image) {
	w.Draw(screen)
}

func (w *Wall) RenderForeground(screen *ebiten.Image) {
	w.Draw(screen)
}
