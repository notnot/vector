// vec3.go, jpad 2022

package vec3

import (
	"fmt"
)

type Vec3 [3]int32

func (v Vec3) String() string {
	return fmt.Sprintf("(%d, %d, %d)", v[0], v[1], v[2])
}
