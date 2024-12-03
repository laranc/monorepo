package graphics3d

import (
	"fmt"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Renderer3D struct {
	window  *sdl.Window
	context sdl.GLContext
}

func MakeRenderer3D(title string, width, height int32, glMajor, glMinor int) (Renderer3D, error) {
	fmt.Println("Initializing...")
	var err error
	var window *sdl.Window
	var context sdl.GLContext

	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return Renderer3D{}, err
	}

	window, err = sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_OPENGL)
	if err != nil {
		return Renderer3D{}, err
	}

	err = sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, glMajor)
	if err != nil {
		return Renderer3D{}, err
	}
	err = sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, glMinor)
	if err != nil {
		return Renderer3D{}, err
	}
	err = sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE)
	if err != nil {
		return Renderer3D{}, err
	}

	context, err = window.GLCreateContext()
	if err != nil {
		return Renderer3D{}, err
	}

	err = gl.Init()
	if err != nil {
		return Renderer3D{}, err
	}

	err = img.Init(img.INIT_PNG)
	if err != nil {
		return Renderer3D{}, err
	}

	renderer := Renderer3D{
		window:  window,
		context: context,
	}
	fmt.Println("Intialization complete")
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL Version", version)
	return renderer, nil
}

func (r *Renderer3D) Destroy() {
	fmt.Println("Closing...")
	sdl.Quit()
	r.window.Destroy()
	sdl.GLDeleteContext(r.context)
	img.Quit()
}

func (r *Renderer3D) RenderBegin() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT | gl.STENCIL_BUFFER_BIT)
}

func (r *Renderer3D) RenderEnd() {
	r.window.GLSwap()
	sdl.Delay(16)
}

func (r *Renderer3D) GetWindow() *sdl.Window {
	return r.window
}

func (r *Renderer3D) GetContext() sdl.GLContext {
	return r.context
}
