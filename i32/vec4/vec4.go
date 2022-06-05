// vec4.go, jpad 2022

package vec4

import (
	"fmt"
)

// Vec4 is a 4-dimensional int32 vector,
type Vec4 [4]int32

func NewVec4() *Vec4 {
	return &Vec4{}
}

func (v Vec4) String() string {
	return fmt.Sprintf("(%d, %d, %d, %d)", v[0], v[1], v[2], v[3])
}
