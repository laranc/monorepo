package builtin

import (
	"github.com/go-gl/mathgl/mgl32"
)

type AABB struct {
	Position mgl32.Vec3
	HalfSize mgl32.Vec3
}
