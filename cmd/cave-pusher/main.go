package main

import (
	"github.com/DmitriiTrifonov/cave-pusher/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
	input "github.com/quasilyte/ebitengine-input"
	"github.com/quasilyte/gmath"
	"log"
)

const (
	TileNoTile int = iota
	TileFloor
	TileWallHorizontalContUpLeft
	TileWallHorizontalContUp
	TileWallHorizontalContUpRight
	TileWallHorizontalContDown
	TileWallVerticalContLeftUp
	TileWallVerticalContLeft
	TileWallVerticalContLeftDown
	TileWallVerticalContRightUp
	TileWallVerticalContRight
	TileWallVerticalContRightDown
)

func main() {
	playerSprite, err := game.NewSprite("assets/exported/man/man.json")
	if err != nil {
		log.Fatal(err)
	}

	tileSet, err := game.NewSprite("assets/exported/tileset/tileset.json")
	if err != nil {
		log.Fatal(err)
	}

	m := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{6, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 9},
		{7, 1, 0, 1, 1, 1, 0, 0, 1, 1, 0, 0, 1, 10},
		{7, 0, 1, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 10},
		{7, 0, 1, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 10},
		{7, 0, 1, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 10},
		{8, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 11},
	}

	spriteTileMap := map[int]*game.TilePrefab{
		TileNoTile: {
			TileNum: 0,
			Sprite:  tileSet,
		},
		TileFloor: {
			TileNum: 1,
			Sprite:  tileSet,
		},
		TileWallHorizontalContUpLeft: {
			TileNum: 2,
			Sprite:  tileSet,
		},
		TileWallHorizontalContUp: {
			TileNum: 3,
			Sprite:  tileSet,
		},
		TileWallHorizontalContUpRight: {
			TileNum: 4,
			Sprite:  tileSet,
		},
		TileWallHorizontalContDown: {
			TileNum: 5,
			Sprite:  tileSet,
		},
		TileWallVerticalContLeftDown: {
			TileNum: 6,
			Sprite:  tileSet,
		},
		TileWallVerticalContLeft: {
			TileNum: 7,
			Sprite:  tileSet,
		},
		TileWallVerticalContLeftUp: {
			TileNum: 8,
			Sprite:  tileSet,
		},
		TileWallVerticalContRightDown: {
			TileNum: 9,
			Sprite:  tileSet,
		},
		TileWallVerticalContRight: {
			TileNum: 10,
			Sprite:  tileSet,
		},
		TileWallVerticalContRightUp: {
			TileNum: 11,
			Sprite:  tileSet,
		},
	}

	tileMapBackground, err := game.NewTileMap(24, spriteTileMap, m)
	if err != nil {
		log.Fatal(err)
	}

	g := &game.Game{}
	g.Debug = true
	game.SetBackground(g, tileMapBackground)
	//game.SetForeground(g, tileMapForeground)
	g.Controller = game.Controller{InputSystem: &input.System{}}
	g.Controller.InputSystem.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})

	keymap := input.Keymap{
		game.ActionMoveLeft:  {input.KeyGamepadLeft, input.KeyLeft, input.KeyA},
		game.ActionMoveRight: {input.KeyGamepadRight, input.KeyRight, input.KeyD},
		game.ActionMoveUp:    {input.KeyGamepadUp, input.KeyUp, input.KeyW},
		game.ActionMoveDown:  {input.KeyGamepadDown, input.KeyDown, input.KeyS},
		game.ActionExit:      {input.KeyEscape},
	}

	g.Controller.InputHandler = g.Controller.InputSystem.NewHandler(0, keymap)

	g.Player = &game.Player{
		Input: g.Controller.InputHandler,
		Object: &game.Object{
			Sprite:   playerSprite,
			Pos:      &gmath.Vec{X: 0, Y: 0},
			IsStatic: false,
		},
		Speed: 120,
	}

	err = g.Player.Object.Sprite.SetAnimTag("idle")
	if err != nil {
		log.Fatal(err)
	}
	g.Player.Object.Sprite.AnimPlayer.PlaySpeed = 0

	ebiten.SetWindowSize(960, 540)
	ebiten.SetWindowTitle("Walls Pusher")
	if err = ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
