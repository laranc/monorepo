package global

const usedKeybinds = 5

type GlobalState struct {
	Time struct {
		Delta, Now, Last                 float64
		FrameLast, FrameDelay, FrameTime float64
		FrameRate, FrameCount            uint64
	}
	Keybinds map[int]uint8
}

var State GlobalState
