package game

import (
	"github.com/quasilyte/gmath"
	"log"
)

type Collider struct {
	Vec      gmath.Vec
	StartPos gmath.Vec
	Height   float64
	Width    float64
}

func (c *Collider) HasCollided(other *Collider) bool {
	if other == nil {
		return false
	}

	if c.Vec.X < other.Vec.X+other.Width &&
		c.Vec.X+other.Width > other.Vec.X &&
		c.Vec.Y < other.Vec.Y+other.Height &&
		c.Vec.Y+c.Height > other.Vec.Y {
		log.Println("collided")
		return true
	}
	return false
}

func (c *Collider) Update(objectPos *gmath.Vec) {
	c.Vec = objectPos.Add(c.StartPos)
}
