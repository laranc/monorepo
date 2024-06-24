package graphics2d

import "github.com/veandco/go-sdl2/sdl"

type Renderer2D struct {
	renderer                     *sdl.Renderer
	shaderProgram                uint32
	vaoQuad, vboQuad, eboQuad    uint32
	vaoLine, vboLine, eboLine    uint32
	vaoBatch, vboBatch, eboBatch uint32
}

func NewRenderer2D(renderer *sdl.Renderer) *Renderer2D {
	return &Renderer2D{
		renderer: renderer,
	}
}

func (r *Renderer2D) Draw() {

}
