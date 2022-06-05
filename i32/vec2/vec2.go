// vec2.go, jpad 2022

package vec2

import (
	"fmt"
)

type Vec2 [2]int32

func (v Vec2) String() string {
	return fmt.Sprintf("(%d, %d)", v[0], v[1])
}
