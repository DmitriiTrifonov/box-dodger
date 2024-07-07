package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gmath"
	"image"
)

const (
	tilesetAnimTag = "tilemap"
)

type TilePrefab struct {
	TileNum     int
	HasCollider bool
	Sprite      *Sprite
}

type Tile struct {
	Object   *Object
	Collider *Collider
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

	var collider *Collider

	if prefab.HasCollider {
		collider = &Collider{
			StartPos: gmath.Vec{0, 0},
			Height:   24,
			Width:    24,
		}

		collider.Update(&pos)
	}

	return &Tile{
		Object: &Object{
			Sprite:       cloned,
			Pos:          &pos,
			IsStatic:     true,
			StartAnimIdx: prefab.TileNum,
		},
		Collider: collider,
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
