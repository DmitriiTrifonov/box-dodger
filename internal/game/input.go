package game

import input "github.com/quasilyte/ebitengine-input"

const (
	ActionExit input.Action = iota
	ActionRestart
	ActionMoveLeft
	ActionMoveRight
	ActionMoveUp
	ActionMoveDown
	ActionTap
	ActionDrag
)
