// vec3.go, jpad 2022

package vec3

import (
	"fmt"
)

type Vec3 [3]float32

func NewVec3() *Vec3 {
	return &Vec3{}
}

func (v Vec3) String() string {
	return fmt.Sprintf("(%.3f, %.3f, %.3f)", v[0], v[1], v[2])
}
