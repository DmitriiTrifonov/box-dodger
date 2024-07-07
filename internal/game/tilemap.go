package game

import "github.com/hajimehoshi/ebiten/v2"

type TileMap struct {
	GridSize int
	Tiles    [][]*Tile
}

func NewTileMap(gridSize int, prefabMap map[int]*TilePrefab, m [][]int) (*TileMap, error) {
	tiles := make([][]*Tile, len(m))

	var err error

	for y, yRow := range m {
		xRow := make([]*Tile, len(yRow))
		for x, val := range yRow {
			prefab, ok := prefabMap[val]
			if !ok {
				return nil, ErrNoPrefab
			}

			if xRow[x], err = NewTile(x, y, gridSize, prefab); err != nil {
				return nil, err
			}

		}
		tiles[y] = xRow
	}

	return &TileMap{
		GridSize: gridSize,
		Tiles:    tiles,
	}, nil
}

func (tm *TileMap) RenderBackground(screen *ebiten.Image) {
	for _, hRow := range tm.Tiles {
		for _, tile := range hRow {
			tile.RenderBackground(screen)
		}
	}
}

func (tm *TileMap) RenderForeground(screen *ebiten.Image) {
	for _, hRow := range tm.Tiles {
		for _, tile := range hRow {
			tile.RenderForeground(screen)
		}
	}
}
