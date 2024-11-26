package graphics2d

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var (
	Black = sdl.Color{R: 0, G: 0, B: 0, A: 255}
	White = sdl.Color{R: 255, G: 255, B: 255, A: 255}
	Red   = sdl.Color{R: 255, G: 0, B: 0, A: 255}
	Green = sdl.Color{R: 0, G: 255, B: 0, A: 255}
	Blue  = sdl.Color{R: 0, G: 0, B: 255, A: 255}
)

type Renderer2D struct {
	window     *sdl.Window
	renderer   *sdl.Renderer
	font       *ttf.Font
	background sdl.Color
	delay      uint32
}

// Constructors and Destructors

func MakeRenderer2D(title string, width, height int32) (Renderer2D, error) {
	fmt.Println("Initializing...")
	var err error
	var window *sdl.Window
	var renderer *sdl.Renderer

	err = sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		return Renderer2D{}, err
	}

	window, err = sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		return Renderer2D{}, err
	}

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return Renderer2D{}, err
	}

	err = img.Init(img.INIT_PNG)
	if err != nil {
		return Renderer2D{}, err
	}

	err = ttf.Init()
	if err != nil {
		return Renderer2D{}, err
	}

	fmt.Println("Initialization complete")
	return Renderer2D{window: window, renderer: renderer, background: Black, delay: 16}, nil
}

func (r *Renderer2D) Destroy() {
	if r.renderer != nil {
		r.renderer.Destroy()
	}
	if r.font != nil {
		r.font.Close()
	}
	if r.window != nil {
		r.window.Destroy()
	}

	sdl.Quit()
	img.Quit()
	ttf.Quit()

}

// Setters

func (r *Renderer2D) SetActiveFont(font *ttf.Font) {
	r.font = font
}

func (r *Renderer2D) SetClearColor(color sdl.Color) {
	r.background = color
}

func (r *Renderer2D) SetDelay(delay uint32) {
	r.delay = delay
}

// Draw Functions

func (r *Renderer2D) RenderBegin() {
	r.renderer.SetDrawColor(r.background.R, r.background.G, r.background.B, r.background.A)
	r.renderer.Clear()
}

func (r *Renderer2D) RenderEnd() {
	r.renderer.Present()
	sdl.Delay(r.delay)
}

func (r *Renderer2D) DrawRect(rect sdl.Rect, color sdl.Color) {
	r.renderer.SetDrawColor(color.R, color.G, color.B, color.A)
	r.renderer.FillRect(&rect)
}

func (r *Renderer2D) DrawRectLines(rect sdl.Rect, color sdl.Color) {
	r.renderer.SetDrawColor(color.R, color.G, color.B, color.A)
	r.renderer.DrawRect(&rect)
}

func (r *Renderer2D) DrawTexture(dst sdl.Rect, src sdl.Rect, texture *sdl.Texture) {
	r.renderer.Copy(texture, &src, &dst)
}

func (r *Renderer2D) DrawText(str string, color sdl.Color, x, y int32) {
	if r.font == nil {
		fmt.Println("Font not loaded")
		return
	}
	text, err := r.font.RenderUTF8Blended(str, color)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer text.Free()
	texture, err := r.renderer.CreateTextureFromSurface(text)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer texture.Destroy()
	r.renderer.Copy(texture, &sdl.Rect{X: 0, Y: 0, W: text.W, H: text.H}, &sdl.Rect{X: x, Y: y, W: text.W, H: text.H})
}

func (r *Renderer2D) DrawAnimation(handler *AnimationHandler, anim *Animation, dst sdl.Rect, color sdl.Color) {
	def := handler.animationDefs[anim.defID]
	frame := &def.frames[anim.currentFrameIndex]
	r.renderer.Copy(def.spriteSheet, &frame.rect, &dst)
}

// Utility Functions

func (r *Renderer2D) CreateTexture(surface *sdl.Surface) (*sdl.Texture, error) {
	return r.renderer.CreateTextureFromSurface(surface)
}
