package input

import (
	"github.com/laranc/monorepo/engine/global"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	KeyStateUnpressed = iota
	KeyStatePressed
	KeyStateHeld
)

func updateKeyState(currentState, keyState uint8) uint8 {
	if currentState != 0 {
		if keyState > 0 {
			return KeyStateHeld
		}
		return KeyStatePressed
	}
	return KeyStateUnpressed
}

func Update() {
	keyboard := sdl.GetKeyboardState()
	for code, state := range global.State.Keybinds {
		global.State.Keybinds[code] = updateKeyState(keyboard[code], state)
	}
}
