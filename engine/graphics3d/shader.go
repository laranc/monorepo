package graphics3d

import (
	"fmt"
	"io"
	"os"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Shader struct {
	id uint32
}

func NewShader(vertPath, fragPath, geoPath string) (*Shader, error) {
	var err error
	var vertexShader, fragmentShader, geometryShader uint32
	shader := new(Shader)

	fmt.Println("Compiling vertex shader...")
	vertexShader, err = shader.loadShader(gl.VERTEX_SHADER, vertPath)
	if err != nil {
		return nil, err
	}
	fmt.Println("Compiling fragment shader...")
	fragmentShader, err = shader.loadShader(gl.FRAGMENT_SHADER, fragPath)
	if err != nil {
		return nil, err
	}
	if geoPath != "" {
		fmt.Println("Compiling geometry shader...")
		geometryShader, err = shader.loadShader(gl.GEOMETRY_SHADER, geoPath)
		if err != nil {
			return nil, err
		}
	}
	shader.linkProgram(vertexShader, fragmentShader, geometryShader)
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)
	gl.DeleteShader(geometryShader)
	return shader, nil
}

func (s *Shader) Destroy() {
	gl.DeleteProgram(s.id)
}

func (s *Shader) loadShader(shaderType uint32, shaderPath string) (uint32, error) {
	shader := gl.CreateShader(shaderType)
	shaderSrc, err := s.loadShaderSrc(shaderPath)
	if err != nil {
		return 0, err
	}
	src, free := gl.Strs(shaderSrc)
	gl.ShaderSource(shader, 1, src, nil)
	free()
	gl.CompileShader(shader)
	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var len int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &len)
		l := make([]byte, len)
		gl.GetShaderInfoLog(shader, len, nil, &l[0])
		panic((string(l)))
	}
	return shader, nil
}

func (s *Shader) loadShaderSrc(shaderPath string) (string, error) {
	fmt.Println("Loading shader from file...")
	var err error
	file, err := os.Open(shaderPath)
	if err != nil {
		fmt.Printf("Failed to open shader file [%s]: %v", shaderPath, err)
		return "", err
	}
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Failed to read shader file [%s]: %v", shaderPath, err)
		return "", err
	}
	file.Close()
	src := string(content) + "\x00"
	fmt.Println("Shader loaded")
	return src, nil
}

func (s *Shader) linkProgram(vertShader, fragShader, geoShader uint32) {
	fmt.Println("Linking shader program...")
	s.id = gl.CreateProgram()
	gl.AttachShader(s.id, vertShader)
	gl.AttachShader(s.id, fragShader)
	if geoShader != 0 {
		gl.AttachShader(s.id, geoShader)
	}
	gl.LinkProgram(s.id)
	var status int32
	gl.GetProgramiv(s.id, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var len int32
		gl.GetProgramiv(s.id, gl.INFO_LOG_LENGTH, &len)
		l := make([]byte, len)
		gl.GetProgramInfoLog(s.id, len, nil, &l[0])
		panic(string(l))
	}
	gl.UseProgram(0)
	fmt.Println("Shader linked")
}

func (s *Shader) Bind() {
	gl.UseProgram(s.id)
}

func (s *Shader) Unbind() {
	gl.UseProgram(0)
}

func (s *Shader) SetInt(value int32, name string) {
	s.Bind()
	gl.Uniform1i(gl.GetUniformLocation(s.id, gl.Str(name+"\x00")), value)
	s.Unbind()
}

func (s *Shader) SetFloat(value float32, name string) {
	s.Bind()
	gl.Uniform1f(gl.GetUniformLocation(s.id, gl.Str(name+"\x00")), value)
	s.Unbind()
}

func (s *Shader) SetVec2(value mgl32.Vec2, name string) {
	s.Bind()
	gl.Uniform2fv(gl.GetUniformLocation(s.id, gl.Str(name+"\x00")), 1, &value[0])
	s.Unbind()
}

func (s *Shader) SetVec3(value mgl32.Vec3, name string) {
	s.Bind()
	gl.Uniform3fv(gl.GetUniformLocation(s.id, gl.Str(name+"\x00")), 1, &value[0])
	s.Unbind()
}

func (s *Shader) SetVec4(value mgl32.Vec4, name string) {
	s.Bind()
	gl.Uniform4fv(gl.GetUniformLocation(s.id, gl.Str(name+"\x00")), 1, &value[0])
	s.Unbind()
}

func (s *Shader) SetMat3(value mgl32.Mat4, name string, transpose bool) {
	s.Bind()
	gl.UniformMatrix3fv(gl.GetUniformLocation(s.id, gl.Str(name+"\x00")), 1, transpose, &value[0])
	s.Unbind()
}

func (s *Shader) SetMat4(value mgl32.Mat4, name string, transpose bool) {
	s.Bind()
	gl.UniformMatrix4fv(gl.GetUniformLocation(s.id, gl.Str(name+"\x00")), 1, transpose, &value[0])
	s.Unbind()
}
