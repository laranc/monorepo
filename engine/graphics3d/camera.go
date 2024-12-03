package graphics3d

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Camera3D struct {
	view, projection           mgl32.Mat4
	position, front, right, up mgl32.Vec3
	orientation                mgl32.Quat
	pitch, yaw, roll           float32
	fov, nearPlane, farPlane   float32
	frameBuffer                mgl32.Vec2
}

func NewCamera(position, up mgl32.Vec3, fov, nearPlane, farPlane float32, frameBuffer mgl32.Vec2) *Camera3D {
	camera := &Camera3D{
		position:    position,
		orientation: mgl32.QuatIdent(),
		up:          up,
		fov:         fov,
		nearPlane:   nearPlane,
		farPlane:    farPlane,
		frameBuffer: frameBuffer,
	}
	camera.Update()
	return camera
}

func (c *Camera3D) Update() {
	c.front = c.orientation.Rotate(mgl32.Vec3{0, 0, -1}).Normalize()
	c.right = c.orientation.Rotate(mgl32.Vec3{1, 0, 0}).Normalize()
	c.up = c.orientation.Rotate(mgl32.Vec3{0, 1, 0}).Normalize()
	c.view = mgl32.LookAtV(c.position, c.position.Add(c.front), c.up)
	c.projection = mgl32.Perspective(mgl32.DegToRad(c.fov), c.frameBuffer[0]/c.frameBuffer[1], c.nearPlane, c.farPlane)
}

func (c *Camera3D) View() mgl32.Mat4 {
	return c.view
}

func (c *Camera3D) Projection() mgl32.Mat4 {
	return c.projection
}

func (c *Camera3D) Position() mgl32.Vec3 {
	return c.position
}

func (c *Camera3D) Front() mgl32.Vec3 {
	return c.front
}

func (c *Camera3D) Right() mgl32.Vec3 {
	return c.right
}

func (c *Camera3D) Move(magnitude mgl32.Vec3) {
	c.position = c.position.Add(magnitude)
}

func (c *Camera3D) Rotate(pitch, yaw, roll float32) {
	c.pitch += pitch
	c.yaw += yaw
	c.roll += roll
	pitchQuat := mgl32.QuatRotate(mgl32.DegToRad(c.pitch), mgl32.Vec3{-1, 0, 0})
	yawQuat := mgl32.QuatRotate(mgl32.DegToRad(c.yaw), mgl32.Vec3{0, -1, 0})
	rollQuat := mgl32.QuatRotate(mgl32.DegToRad(c.roll), mgl32.Vec3{0, 0, -1})
	rotation := yawQuat.Mul(pitchQuat).Mul(rollQuat)
	c.orientation = rotation.Normalize()
}
