package ui

import "github.com/go-gl/mathgl/mgl32"

type Button struct {
	down     bool
	text     string
	position mgl32.Vec2
	size     mgl32.Vec2
}

type ButtonEvent interface {
	Register()
	Trigger(payload bool)
}

func (b *Button) Handle(payload bool) {

}

func (b *Button) Register()

func (b *Button) Trigger(payload bool) {
}
