package game

import (
	"github.com/quasilyte/gmath"
)

var (
	collisionsMap = map[string][]*Collider{}
)

func AddCollisionToTag(tag string, col *Collider) {
	if _, ok := collisionsMap[tag]; !ok {
		collisionsMap[tag] = []*Collider{col}

		return
	}

	collisionsMap[tag] = append(collisionsMap[tag], col)
}

func GetCollisionsFromTag(tag string) []*Collider {
	return collisionsMap[tag]
}

type Collider struct {
	Vec      gmath.Vec
	StartPos gmath.Vec
	Height   float64
	Width    float64
	Tag      string
}

func (c *Collider) CheckCollisionsWithTag(tag string) bool {
	for _, col := range collisionsMap[tag] {
		ok := c.HasCollided(col)
		if ok {
			return true
		}
	}

	return false
}

func (c *Collider) HasCollided(other *Collider) bool {
	if other == nil {
		return false
	}

	if c.Vec.X < other.Vec.X+other.Width &&
		c.Vec.X+other.Width > other.Vec.X &&
		c.Vec.Y < other.Vec.Y+other.Height &&
		c.Vec.Y+c.Height > other.Vec.Y {
		return true
	}
	return false
}

func (c *Collider) Update(objectPos gmath.Vec) {
	c.Vec = objectPos.Add(c.StartPos)
}
