// vec3.go, jpad 2022

package vec3

import (
	"fmt"
)

// Vec3 is a 3-dimensional int32 vector.
type Vec3 [3]int32

func NewVec3() *Vec3 {
	return &Vec3{}
}

func (v Vec3) String() string {
	return fmt.Sprintf("(%d, %d, %d)", v[0], v[1], v[2])
}
