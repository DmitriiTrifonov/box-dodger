package main

import (
	"github.com/DmitriiTrifonov/cave-pusher/internal/game"
	"github.com/DmitriiTrifonov/cave-pusher/internal/player"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	input "github.com/quasilyte/ebitengine-input"
	"github.com/quasilyte/gmath"
	"github.com/solarlune/goaseprite"
	"log"
)

func main() {
	sprite := goaseprite.Open("assets/exported/man/man.json")

	img, _, err := ebitenutil.NewImageFromFile(sprite.ImagePath)
	if err != nil {
		log.Fatal(err)
	}

	g := &game.Game{}
	g.Debug = true
	g.InputSystem.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})
	keymap := input.Keymap{
		player.ActionMoveLeft:  {input.KeyGamepadLeft, input.KeyLeft, input.KeyA},
		player.ActionMoveRight: {input.KeyGamepadRight, input.KeyRight, input.KeyD},
		player.ActionMoveUp:    {input.KeyGamepadUp, input.KeyUp, input.KeyW},
		player.ActionMoveDown:  {input.KeyGamepadDown, input.KeyDown, input.KeyS},
	}

	g.Player = &player.Player{
		Input:  g.InputSystem.NewHandler(0, keymap),
		Pos:    &gmath.Vec{X: 0, Y: 0},
		Sprite: sprite,
		Anim:   sprite.CreatePlayer(),
		Img:    img,
		Speed:  120,
	}

	ebiten.SetWindowSize(960, 540)
	ebiten.SetWindowTitle("Cave Pusher")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
