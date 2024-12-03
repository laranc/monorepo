package graphics3d

import "github.com/go-gl/mathgl/mgl32"

type Model struct {
	position         mgl32.Vec3
	material         *Material
	overrideDiffuse  *Texture
	overrideSpecular *Texture
	meshes           []*Mesh
	shader           *Shader
}

func NewModel(position mgl32.Vec3, material *Material, diffuse, specular *Texture, meshes []*Mesh, shader *Shader) *Model {
	model := &Model{
		position:         position,
		material:         material,
		overrideDiffuse:  diffuse,
		overrideSpecular: specular,
		meshes:           meshes,
		shader:           shader,
	}
	return model
}

func (m *Model) Render() {
	m.material.SendToShader(m.shader)
	m.overrideDiffuse.Bind()
	m.overrideSpecular.Bind()
	for _, mesh := range m.meshes {
		mesh.Render(m.shader)
	}
	m.overrideDiffuse.Unbind()
	m.overrideSpecular.Unbind()
}
