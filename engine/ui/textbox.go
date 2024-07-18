package ui

import "github.com/go-gl/mathgl/mgl32"

type TextBox struct {
	text     string
	position mgl32.Vec2
	size     mgl32.Vec2
}

func (tb *TextBox) SetPos(pos mgl32.Vec2) {
	tb.position = pos
}

func (tb *TextBox) GetPos() mgl32.Vec2 {
	return tb.position
}
