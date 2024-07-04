package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	input "github.com/quasilyte/ebitengine-input"
)

type Controller struct {
	InputSystem  *input.System
	InputHandler *input.Handler
}

func (c *Controller) Update() error {
	c.InputSystem.Update()
	if c.InputHandler.ActionIsPressed(ActionExit) {
		return ebiten.Termination
	}
	return nil
}
