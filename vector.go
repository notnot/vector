// vector.go, jpad 2022

package main

import (
	"fmt"

	vec2f "notnot.org/vector/f32/vec2"
	vec3f "notnot.org/vector/f32/vec3"
	vec2i "notnot.org/vector/i32/vec2"
	vec3i "notnot.org/vector/i32/vec3"
)

// main() does some vector manipulations
func main() {
	fmt.Printf("playing with some vectors...\n")

	v2i := vec2i.Vec2{}
	fmt.Printf("%s\n", v2i.String())

	v2f := vec2f.Vec2{}
	fmt.Printf("%s\n", v2f.String())

	v3i := vec3i.Vec3{}
	fmt.Printf("%s\n", v3i.String())

	v3f := vec3f.Vec3{}
	fmt.Printf("%s\n", v3f.String())
}
