package game

import (
	"fmt"
	"github.com/DmitriiTrifonov/cave-pusher/internal/player"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	input "github.com/quasilyte/ebitengine-input"
)

type Game struct {
	Debug       bool
	InputSystem input.System
	Player      *player.Player
}

func (g *Game) Update() error {
	g.Player.Anim.Update(float32(1.0 / 60.0))
	g.Player.Update(ebiten.ActualTPS())
	g.InputSystem.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.Debug {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("X: %v, Y: %v", g.Player.Pos.X, g.Player.Pos.Y))
	}
	g.Player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 180
}
