package main

import (
	"github.com/DmitriiTrifonov/cave-pusher/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	input "github.com/quasilyte/ebitengine-input"
	"github.com/quasilyte/gmath"
	"github.com/solarlune/goaseprite"
	"log"
)

func main() {
	playerSprFile := goaseprite.Open("assets/exported/man/man.json")

	playerImg, _, err := ebitenutil.NewImageFromFile(playerSprFile.ImagePath)
	if err != nil {
		log.Fatal(err)
	}

	wallSprFile := goaseprite.Open("assets/exported/wall/wall.json")
	wallImg, _, err := ebitenutil.NewImageFromFile(wallSprFile.ImagePath)
	if err != nil {
		log.Fatal(err)
	}

	bgWall := &game.Wall{
		Pos:    &gmath.Vec{X: 120, Y: 100},
		Sprite: wallSprFile,
		Anim:   wallSprFile.CreatePlayer(),
		Img:    wallImg,
	}

	fgWall := &game.Wall{
		Pos:    &gmath.Vec{X: 144, Y: 100},
		Sprite: wallSprFile,
		Anim:   wallSprFile.CreatePlayer(),
		Img:    wallImg,
	}

	g := &game.Game{}
	g.Debug = true
	g.BackgroundRender = []game.BackgroundRenderer{bgWall}
	g.ForegroundRenderer = []game.ForegroundRenderer{fgWall}
	g.InputSystem.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})
	keymap := input.Keymap{
		game.ActionMoveLeft:  {input.KeyGamepadLeft, input.KeyLeft, input.KeyA},
		game.ActionMoveRight: {input.KeyGamepadRight, input.KeyRight, input.KeyD},
		game.ActionMoveUp:    {input.KeyGamepadUp, input.KeyUp, input.KeyW},
		game.ActionMoveDown:  {input.KeyGamepadDown, input.KeyDown, input.KeyS},
	}

	g.Player = &game.Player{
		Input:  g.InputSystem.NewHandler(0, keymap),
		Pos:    &gmath.Vec{X: 0, Y: 0},
		Sprite: playerSprFile,
		Anim:   playerSprFile.CreatePlayer(),
		Img:    playerImg,
		Speed:  120,
	}

	ebiten.SetWindowSize(960, 540)
	ebiten.SetWindowTitle("Walls Pusher")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
