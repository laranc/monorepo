package graphics2d

import "github.com/veandco/go-sdl2/sdl"

const MaxFrames = 16

type AnimationFrame struct {
	duration float32
	rect     sdl.Rect
}

type AnimationDef struct {
	spriteSheet *sdl.Texture
	frames      [MaxFrames]AnimationFrame
	frameCount  uint
}

type Animation struct {
	defID                         uint64
	currentFrameTime              float32
	currentFrameIndex             uint
	doesLoop, isActive, isFlipped bool
}

type AnimationHandler struct {
	animationDefs []*AnimationDef
	animations    []*Animation
}

func MakeAnimationHandler() AnimationHandler {
	return AnimationHandler{
		animationDefs: make([]*AnimationDef, 0),
		animations:    make([]*Animation, 0),
	}
}

func (h *AnimationHandler) CreateAnimationDef(spriteSheet *sdl.Texture, duration float32, frameCount uint, frameRect sdl.Rect) uint64 {
	def := &AnimationDef{
		spriteSheet: spriteSheet,
		frameCount:  frameCount,
	}
	for i := range frameCount {
		def.frames[i] = AnimationFrame{
			duration: duration,
			rect:     frameRect,
		}
	}
	id := uint64(len(h.animationDefs))
	h.animationDefs = append(h.animationDefs, def)
	return id
}

func (h *AnimationHandler) CreateAnimation(animationDefID uint64, doesLoop bool) uint64 {
	id := uint64(len(h.animations))
	def := h.animationDefs[animationDefID]
	if def == nil {
		return id
	}
	for i, anim := range h.animations {
		if !anim.isActive {
			id = uint64(i)
			break
		}
	}
	if id == uint64(len(h.animations)) {
		h.animations = append(h.animations, new(Animation))
	}
	anim := h.animations[id]
	*anim = Animation{
		defID:    animationDefID,
		doesLoop: doesLoop,
		isActive: true,
	}
	return id
}

func (h *AnimationHandler) DestroyAnimation(id uint64) {
	anim := h.animations[id]
	anim.isActive = false
}

func (h *AnimationHandler) AnimationUpdate(dt float32) {
	for _, anim := range h.animations {
		def := h.animationDefs[anim.defID]
		anim.currentFrameTime -= dt
		if anim.currentFrameTime <= 0 {
			anim.currentFrameIndex += 1
			if anim.currentFrameIndex == def.frameCount {
				if anim.doesLoop {
					anim.currentFrameIndex = 0
				} else {
					anim.currentFrameIndex -= 1
				}
			}
			anim.currentFrameTime = def.frames[anim.currentFrameIndex].duration
		}
	}
}
