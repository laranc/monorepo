package graphics3d

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Vertex struct {
	Position mgl32.Vec3
	Color    mgl32.Vec3
	Texcoord mgl32.Vec2
	Normal   mgl32.Vec3
}
