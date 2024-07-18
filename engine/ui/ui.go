package ui

import "github.com/go-gl/mathgl/mgl32"

type Widget interface {
	SetPos(pos mgl32.Vec2)
	GetPos() mgl32.Vec2
}

type UI struct {
	boxes []Box
}

func (ui *UI) Begin() {

}

func (ui *UI) End() {

}

func (ui *UI) CreateBox(pos, size mgl32.Vec2) Box {
	box := Box{
		position: pos,
		size:     size,
	}
	ui.boxes = append(ui.boxes, box)
	return box
}
