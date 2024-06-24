package graphics3d

import "github.com/go-gl/mathgl/mgl32"

type Material struct {
	ambient, diffuse, specular      mgl32.Vec3
	diffuseTexture, specularTexture *Texture
}

func NewMaterial(ambient, diffuse, specular mgl32.Vec3, diffuseTexture, specularTexture *Texture) *Material {
	return &Material{
		ambient:         ambient,
		diffuse:         diffuse,
		specular:        specular,
		diffuseTexture:  diffuseTexture,
		specularTexture: specularTexture,
	}
}

func (m *Material) Destroy() {
	if m.diffuseTexture != nil {
		m.diffuseTexture.Destroy()
	}
	if m.specularTexture != nil {
		m.specularTexture.Destroy()
	}
}

func (m *Material) SendToShader(shader *Shader) {
	shader.SetVec3(m.ambient, "material.ambient")
	shader.SetVec3(m.diffuse, "material.diffuse")
	shader.SetVec3(m.specular, "material.specular")
	if m.diffuseTexture != nil {
		shader.SetInt(m.diffuseTexture.Unit(), "material.diffuse_texture")
	}
	if m.specularTexture != nil {
		shader.SetInt(m.specularTexture.Unit(), "material.specular_texture")
	}
}
