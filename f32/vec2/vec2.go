// vec2.go, jpad 2022

package vec2

import (
	"fmt"
)

// Vec2 is a 2-dimensional float32 vector,
type Vec2 [2]float32

func NewVec2() *Vec2 {
	return &Vec2{}
}

func (v Vec2) String() string {
	return fmt.Sprintf("(%.3f, %.3f)", v[0], v[1])
}
