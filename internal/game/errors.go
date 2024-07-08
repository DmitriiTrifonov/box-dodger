package game

import "errors"

var (
	ErrNoPrefab    = errors.New("no prefab")
	ErrRestartGame = errors.New("restart game")
)
