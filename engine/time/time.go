package time

import (
	"github.com/laranc/monorepo/engine/global"
	"github.com/veandco/go-sdl2/sdl"
)

func Init(frameRate uint64) {
	global.State.Time.FrameRate = frameRate
	global.State.Time.FrameDelay = 1000.0 / float64(frameRate)
}

func Update() {
	global.State.Time.Now = float64(sdl.GetTicks64())
	global.State.Time.Delta = (global.State.Time.Now - global.State.Time.Last) / 1000.0
	global.State.Time.Last = global.State.Time.Now
	global.State.Time.FrameCount++
	if global.State.Time.Now-global.State.Time.FrameLast >= 1000.0 {
		global.State.Time.FrameRate = global.State.Time.FrameCount
		global.State.Time.FrameCount = 0
		global.State.Time.FrameLast = global.State.Time.Now
	}
}

func UpdateLate() {
	global.State.Time.FrameTime = float64(sdl.GetTicks64()) - global.State.Time.Now
	if global.State.Time.FrameDelay > global.State.Time.FrameTime {
		sdl.Delay(uint32(global.State.Time.FrameDelay - global.State.Time.FrameTime))
	}
}
