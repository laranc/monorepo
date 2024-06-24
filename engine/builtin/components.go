package builtin

import (
	"reflect"

	"github.com/go-gl/mathgl/mgl32"
)

type Transform struct {
	Translation mgl32.Vec3
	Rotation    mgl32.Vec3
	Scale       mgl32.Vec3
	reflectType reflect.Type
}

func MakeTransform(translation, rotation, scale mgl32.Vec3) Transform {
	t := Transform{
		Translation: translation,
		Rotation:    rotation,
		Scale:       scale,
	}
	t.reflectType = reflect.TypeOf(t)
	return t
}

func (t Transform) Type() reflect.Type {
	return t.reflectType
}

type Body struct {
	AABB         AABB
	Velocity     mgl32.Vec3
	Acceleration mgl32.Vec3
	reflectType  reflect.Type
}

func MakeBody(aabb AABB, velocity, acceleration mgl32.Vec3) Body {
	b := Body{
		AABB:         aabb,
		Velocity:     velocity,
		Acceleration: acceleration,
	}
	b.reflectType = reflect.TypeOf(b)
	return b
}

func (b *Body) Type() reflect.Type {
	return b.reflectType
}
