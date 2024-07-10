package game

import (
	"errors"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	collisionTagWalls = "walls"
)

type Game struct {
	Debug              bool
	Controller         Controller
	Player             *Player
	TileMap            *TileMap
	BackgroundRender   []BackgroundRenderer
	ForegroundRenderer []ForegroundRenderer
}

func New(controller Controller, player *Player, tileMap TileMap) *Game {
	g := &Game{}

	g.Debug = true
	g.Controller = controller

	return g
}

type BackgroundRenderer interface {
	RenderBackground(screen *ebiten.Image)
}

type ForegroundRenderer interface {
	RenderForeground(screen *ebiten.Image)
}

func (g *Game) Update() error {
	g.Player.Object.Sprite.AnimPlayer.Update(float32(1.0 / 60.0))
	g.Player.Update(ebiten.TPS())
	g.Player.HasCollided = g.CheckCollisions()
	err := g.Controller.Update()
	if err != nil {
		return err
	}

	return nil
}

func (g *Game) CheckCollisions() bool {
	for _, hRow := range g.TileMap.Tiles {
		for _, tile := range hRow {
			ok := g.Player.Collider.HasCollided(tile.Collider)
			if ok {
				return true
			}
		}
	}

	return false
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, gr := range g.BackgroundRender {
		gr.RenderBackground(screen)
	}
	g.Player.Draw(screen)
	for _, fr := range g.ForegroundRenderer {
		fr.RenderForeground(screen)
	}

	if g.Debug {
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Collided: %v",
			g.Player.HasCollided), 0, 0)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 180
}

func SetBackground[T BackgroundRenderer](game *Game, objects ...T) {
	bg := make([]BackgroundRenderer, len(objects))
	for i, obj := range objects {
		bg[i] = obj
	}

	game.BackgroundRender = bg
}

func SetForeground[T ForegroundRenderer](game *Game, objects ...T) {
	fg := make([]ForegroundRenderer, len(objects))
	for i, obj := range objects {
		fg[i] = obj
	}

	game.ForegroundRenderer = fg
}

func (g *Game) Run() error {
	err := ebiten.RunGame(g)
	switch {
	case errors.Is(err, ErrRestartGame):
		return g.Run()
	default:
		return err
	}
}
