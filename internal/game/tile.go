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
	TileNum  int
	Collider *Collider
	Sprite   *Sprite
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

	if prefab.Collider != nil {
		collider = &Collider{
			StartPos: prefab.Collider.StartPos,
			Height:   prefab.Collider.Height,
			Width:    prefab.Collider.Width,
			Tag:      collisionTagWalls,
		}
		collider.Update(pos)
	}

	AddCollisionToTag(collisionTagWalls, collider)

	return &Tile{
		Object: &Object{
			Sprite:       cloned,
			Pos:          pos,
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
