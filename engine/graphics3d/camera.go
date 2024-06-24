package graphics3d

import "github.com/go-gl/mathgl/mgl32"

type Camera3D struct {
	Model      mgl32.Mat4
	View       mgl32.Mat4
	Projection mgl32.Mat4
	Position   mgl32.Vec3
	Up         mgl32.Vec3
	Front      mgl32.Vec3
	FOV        float32
}

func NewCamera(model mgl32.Mat4, view, projection mgl32.Mat4, position, up, front mgl32.Vec3, fov float32) *Camera3D {
	return &Camera3D{
		Model:      model,
		View:       view,
		Projection: projection,
		Position:   position,
		Up:         up,
		Front:      front,
		FOV:        fov,
	}
}