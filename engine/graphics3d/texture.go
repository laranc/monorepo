package graphics3d

import (
	"fmt"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Texture struct {
	id     uint32
	target uint32
	unit   uint32
	rect   sdl.Rect
	UV1    mgl32.Vec2 // READ ONLY
	UV2    mgl32.Vec2 // READ ONLY
}

func NewTexture(rect sdl.Rect, imgPath string, target, unit uint32) (*Texture, error) {
	fmt.Println("Making texture...")
	var uv1, uv2, imgSize mgl32.Vec2
	var id uint32

	image, err := img.Load(imgPath)
	if err != nil {
		fmt.Printf("Error loading image file: %v\n", err)
		return nil, err
	}
	defer image.Free()

	imgSize = mgl32.Vec2{float32(image.W), float32(image.H)}
	if !rect.Empty() {
		fmt.Println("Texture is a texture atlas")
		uv1 = mgl32.Vec2{float32(rect.X*rect.W) / imgSize[0], float32(rect.Y*rect.H) / imgSize[1]}
		uv2 = mgl32.Vec2{float32(rect.W+(rect.X*rect.W)) / imgSize[0], float32(rect.H+(rect.Y*rect.H)) / imgSize[1]}
	} else {
		fmt.Println("Texture is a normal texture")
		uv1 = mgl32.Vec2{0.0, 0.0}
		uv2 = imgSize
	}

	gl.GenTextures(1, &id)
	gl.BindTexture(gl.TEXTURE_2D, id)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)

	var format uint32
	switch image.Format.Format {
	case uint32(sdl.PIXELFORMAT_RGBA32):
		format = gl.RGBA
	case sdl.PIXELFORMAT_RGB24:
		format = gl.RGB
	default:
		panic("Unsupported surface format")
	}

	convertedSurface, err := image.ConvertFormat(image.Format.Format, 0)
	if err != nil {
		fmt.Printf("Error converting surface format: %v\n", err)
		return nil, err
	}
	defer convertedSurface.Free()

	gl.TexImage2D(gl.TEXTURE_2D, 0, int32(format), convertedSurface.W, convertedSurface.H, 0, format, gl.UNSIGNED_BYTE, gl.Ptr(convertedSurface.Pixels()))
	gl.GenerateMipmap(gl.TEXTURE_2D)

	gl.ActiveTexture(0)
	gl.BindTexture(gl.TEXTURE_2D, 0)
	texture := &Texture{
		id:     id,
		target: target,
		unit:   unit,
		rect:   rect,
		UV1:    uv1,
		UV2:    uv2,
	}
	fmt.Println("Texture created successfully")
	return texture, nil
}

func (t *Texture) Destroy() {
	gl.DeleteTextures(1, &t.id)
}

func (t *Texture) ID() uint32 {
	return t.id
}

func (t *Texture) Unit() int32 {
	return int32(t.unit)
}

func (t *Texture) Bind() {
	gl.ActiveTexture(gl.TEXTURE0 + t.unit)
	gl.BindTexture(t.target, t.id)
}

func (t *Texture) Unbind() {
	gl.ActiveTexture(0)
	gl.BindTexture(t.target, 0)
}
