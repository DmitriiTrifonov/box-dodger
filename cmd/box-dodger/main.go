package main

import (
	"github.com/DmitriiTrifonov/cave-pusher/internal/game"
	input "github.com/quasilyte/ebitengine-input"
	"github.com/quasilyte/gmath"
	"log"
	"time"
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

	boxSprite, err := game.NewSprite("assets/exported/box/box.json")
	if err != nil {
		log.Fatal(err)
	}

	m := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{6, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 9},
		{7, 1, 0, 1, 1, 1, 0, 0, 1, 1, 0, 0, 10},
		{7, 0, 1, 1, 0, 0, 1, 0, 0, 0, 1, 0, 10},
		{7, 0, 1, 1, 0, 0, 0, 1, 0, 0, 0, 1, 10},
		{7, 0, 1, 1, 0, 0, 1, 0, 0, 0, 1, 0, 10},
		{8, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 11},
	}

	colliderFull := &game.Collider{
		StartPos: gmath.Vec{0, 0},
		Height:   24,
		Width:    24,
	}

	colliderDown := &game.Collider{
		StartPos: gmath.Vec{0, 16},
		Height:   24 - 16,
		Width:    24,
	}

	colliderRight := &game.Collider{
		StartPos: gmath.Vec{0, 0},
		Height:   24,
		Width:    24 - 16,
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
			TileNum:  2,
			Sprite:   tileSet,
			Collider: colliderFull,
		},
		TileWallHorizontalContUp: {
			TileNum:  3,
			Sprite:   tileSet,
			Collider: colliderFull,
		},
		TileWallHorizontalContUpRight: {
			TileNum:  4,
			Sprite:   tileSet,
			Collider: colliderFull,
		},
		TileWallHorizontalContDown: {
			TileNum:  5,
			Sprite:   tileSet,
			Collider: colliderDown,
		},
		TileWallVerticalContLeftDown: {
			TileNum:  6,
			Sprite:   tileSet,
			Collider: colliderFull,
		},
		TileWallVerticalContLeft: {
			TileNum:  7,
			Sprite:   tileSet,
			Collider: colliderFull,
		},
		TileWallVerticalContLeftUp: {
			TileNum:  8,
			Sprite:   tileSet,
			Collider: colliderFull,
		},
		TileWallVerticalContRightDown: {
			TileNum:  9,
			Sprite:   tileSet,
			Collider: colliderRight,
		},
		TileWallVerticalContRight: {
			TileNum:  10,
			Sprite:   tileSet,
			Collider: colliderRight,
		},
		TileWallVerticalContRightUp: {
			TileNum:  11,
			Sprite:   tileSet,
			Collider: colliderRight,
		},
	}

	tileMapBackground, err := game.NewTileMap(24, spriteTileMap, m)
	if err != nil {
		log.Fatal(err)
	}

	g := &game.Game{}
	g.Debug = true
	game.SetBackground(g, tileMapBackground)
	g.TileMap = tileMapBackground
	g.StartTime = time.Now()
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
		game.ActionRestart:   {input.KeyR},
	}

	g.Controller.InputHandler = g.Controller.InputSystem.NewHandler(0, keymap)

	g.Player = game.NewPlayer(
		120,
		&game.Object{
			Sprite:   playerSprite,
			Pos:      gmath.Vec{X: 48, Y: 48},
			IsStatic: false,
		},
		&game.Collider{
			StartPos: gmath.Vec{6, 16},
			Height:   6,
			Width:    12,
		},
		g.Controller.InputHandler)

	err = g.Player.Object.Sprite.SetAnimTag("idle")
	if err != nil {
		log.Fatal(err)
	}
	g.Player.Object.Sprite.AnimPlayer.PlaySpeed = 0

	g.Box = game.NewBox(
		120,
		&game.Object{
			Sprite:   boxSprite,
			Pos:      gmath.Vec{48.0, 72.0},
			IsStatic: false,
		},
		&game.Collider{
			Height: 16,
			Width:  16,
			Tag:    "box",
		})

	err = g.Box.Object.Sprite.SetAnimTag("")
	if err != nil {
		log.Fatal(err)
	}
	g.Box.Object.Sprite.AnimPlayer.PlaySpeed = 0

	if err = g.Run(); err != nil {
		log.Fatal(err)
	}
}
