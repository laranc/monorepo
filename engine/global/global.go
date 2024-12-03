package global

import "github.com/veandco/go-sdl2/sdl"

var State struct {
	Time struct {
		Delta, Now, Last float32
	}
	Keyboard struct {
		Binds  map[string]sdl.Scancode
		States map[sdl.Scancode]uint8
	}
	Mouse struct {
		X, Y float32
	}
}

func init() {
	State.Keyboard.Binds = make(map[string]sdl.Scancode)
	State.Keyboard.States = make(map[sdl.Scancode]uint8)
}

func UpdateTime() {
	State.Time.Now = float32(sdl.GetTicks64())
	State.Time.Delta = State.Time.Now - State.Time.Last
	State.Time.Last = State.Time.Now
}

func UpdateKeyboard() {
	keys := sdl.GetKeyboardState()
	for _, v := range State.Keyboard.Binds {
		State.Keyboard.States[v] = keys[v]
	}
}

func UpdateMouse() {
	x, y, _ := sdl.GetRelativeMouseState()
	State.Mouse.X = float32(x)
	State.Mouse.Y = float32(y)
}

func UpdateAll() {
	UpdateTime()
	UpdateKeyboard()
	UpdateMouse()
}
