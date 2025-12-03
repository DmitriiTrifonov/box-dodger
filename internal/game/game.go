package game

import (
	"errors"
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	collisionTagWalls  = "walls"
	collisionTagBox    = "box"
	collisionTagPlayer = "player"
)

var debugText string

type Game struct {
	Debug              bool
	Controller         Controller
	Player             *Player
	Box                *Box
	TileMap            *TileMap
	BackgroundRender   []BackgroundRenderer
	ForegroundRenderer []ForegroundRenderer
	IsGameOver         bool
	StartTime          time.Time
	TotalTime          time.Duration
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
	err := g.Controller.Update()
	if err != nil {
		return err
	}

	if !g.IsGameOver {
		g.Player.Object.Sprite.AnimPlayer.Update(float32(1.0 / 60.0))
		g.Player.Update(ebiten.TPS())
		//g.Player.HasCollided = g.CheckCollisions()
		g.Box.Update(ebiten.TPS(), g)
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
	if g.IsGameOver {
		screen.Fill(color.RGBA{R: 0xff, A: 0xff})
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Game Over\nYou survived for: %.1f seconds",
			g.TotalTime.Seconds()), 320/4, 160/4)

	} else {
		for _, gr := range g.BackgroundRender {
			gr.RenderBackground(screen)
		}
		g.Player.Draw(screen)
		g.Player.VirtualJoystick.Draw(screen)
		g.Box.Draw(screen)

		for _, fr := range g.ForegroundRenderer {
			fr.RenderForeground(screen)
		}
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Dodge time: %.1f seconds",
			time.Since(g.StartTime).Seconds()), 16, 8)
	}

	if g.Debug {

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
	ebiten.SetWindowSize(1600, 960)
	ebiten.SetWindowTitle("box dodger")
	err := ebiten.RunGame(g)
	switch {
	case errors.Is(err, ErrRestartGame):
		return g.Run()
	default:
		return err
	}
}
