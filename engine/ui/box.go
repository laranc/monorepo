package ui

import "github.com/go-gl/mathgl/mgl32"

type Box struct {
	position mgl32.Vec2
	size     mgl32.Vec2
	widgets  []Widget
}

func (b *Box) Add(widget Widget) {
	b.widgets = append(b.widgets, widget)
}
