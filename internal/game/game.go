package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Debug              bool
	Controller         Controller
	Player             *Player
	BackgroundRender   []BackgroundRenderer
	ForegroundRenderer []ForegroundRenderer
}

type BackgroundRenderer interface {
	RenderBackground(screen *ebiten.Image)
}

type ForegroundRenderer interface {
	RenderForeground(screen *ebiten.Image)
}

func (g *Game) Update() error {
	g.Player.Anim.Update(float32(1.0 / 60.0))
	g.Player.Update(ebiten.ActualTPS())
	err := g.Controller.Update()
	if err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.Debug {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("X: %v, Y: %v", g.Player.Pos.X, g.Player.Pos.Y))
	}
	for _, fr := range g.BackgroundRender {
		fr.RenderBackground(screen)
	}
	g.Player.Draw(screen)
	for _, fr := range g.ForegroundRenderer {
		fr.RenderForeground(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 180
}
