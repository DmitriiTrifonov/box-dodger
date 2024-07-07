package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gmath"
	"image"
	"log"
)

const (
	tilesetAnimTag = "tilemap"
)

type TilePrefab struct {
	TileNum int
	Sprite  *Sprite
}

type Tile struct {
	Object *Object
}

func NewTile(gridX, gridY, gridSize int, prefab *TilePrefab) (*Tile, error) {
	pos := gmath.VecFromStd(image.Point{
		X: gridX * gridSize,
		Y: gridY * gridSize,
	})

	cloned := prefab.Sprite.Clone()
	err := cloned.SetAnimTag(tilesetAnimTag)
	if err != nil {
		return nil, fmt.Errorf("cannot set tileset anim: %w", err)
	}

	cloned.AnimPlayer.SetFrameIndex(prefab.TileNum)
	cloned.AnimPlayer.PlaySpeed = 0

	log.Println(prefab.Sprite.AnimPlayer.CurrentFrame())

	return &Tile{
		Object: &Object{
			Sprite:       cloned,
			Pos:          &pos,
			IsStatic:     true,
			StartAnimIdx: prefab.TileNum,
		},
	}, nil
}

func (t *Tile) Draw(screen *ebiten.Image) {
	t.Object.Draw(screen)
}

func (t *Tile) RenderBackground(screen *ebiten.Image) {
	t.Draw(screen)
}

func (t *Tile) RenderForeground(screen *ebiten.Image) {
	t.Draw(screen)
}
