package graphics3d

import (
	"unsafe"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Mesh struct {
	vao, vbo, ebo                          uint32
	vertexCount, indexCount, triangleCount int
	position, rotation, scale              mgl32.Vec3
	model                                  mgl32.Mat4
}

func NewMesh(vertices []Vertex, indices []uint32, vertexCount, indexCount, triangleCount int, position, rotation, scale mgl32.Vec3) *Mesh {
	mesh := &Mesh{
		vertexCount:   vertexCount,
		indexCount:    indexCount,
		triangleCount: triangleCount,
		position:      position,
		rotation:      rotation,
		scale:         scale,
	}
	mesh.initGLData(vertices, indices)
	mesh.updateModel()
	return mesh
}

func NewMeshPrimitive(primitive Primitive, position, rotation, scale mgl32.Vec3) *Mesh {
	return NewMesh(primitive.GetVertices(), primitive.GetIndices(), primitive.GetVertexCount(), primitive.GetIndexCount(), primitive.GetTriangleCount(), position, rotation, scale)
}

func (m *Mesh) Destroy() {
	gl.DeleteVertexArrays(1, &m.vao)
	gl.DeleteBuffers(1, &m.vbo)
	gl.DeleteBuffers(1, &m.ebo)
}

func (m *Mesh) initGLData(vertices []Vertex, indices []uint32) {
	gl.GenVertexArrays(1, &m.vao)
	gl.GenBuffers(1, &m.vbo)

	gl.BindVertexArray(m.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, m.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, m.vertexCount*int(unsafe.Sizeof(vertices[0])), gl.Ptr(vertices), gl.STATIC_DRAW)

	if m.indexCount > 0 {
		gl.GenBuffers(1, &m.ebo)
		gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, m.ebo)
		gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, m.triangleCount*3*int(unsafe.Sizeof(indices[0])), gl.Ptr(indices), gl.STATIC_DRAW)
	}

	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, int32(unsafe.Sizeof(vertices[0])), unsafe.Offsetof(vertices[0].Position))
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(1, 3, gl.FLOAT, false, int32(unsafe.Sizeof(vertices[0])), unsafe.Offsetof(vertices[0].Color))
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointerWithOffset(2, 2, gl.FLOAT, false, int32(unsafe.Sizeof(vertices[0])), unsafe.Offsetof(vertices[0].Texcoord))
	gl.EnableVertexAttribArray(2)
	gl.VertexAttribPointerWithOffset(3, 3, gl.FLOAT, false, int32(unsafe.Sizeof(vertices[0])), unsafe.Offsetof(vertices[0].Normal))
	gl.EnableVertexAttribArray(3)

	gl.BindVertexArray(0)
}

func (m *Mesh) SetPosition(position mgl32.Vec3) {
	m.position = position
}

func (m *Mesh) SetRotation(rotation mgl32.Vec3) {
	m.rotation = rotation
}

func (m *Mesh) SetScale(scale mgl32.Vec3) {
	m.scale = scale
}

func (m *Mesh) Move(magnitude mgl32.Vec3) {
	m.position = m.position.Add(magnitude)
}

func (m *Mesh) Rotate(angle mgl32.Vec3) {
	m.rotation = m.rotation.Add(angle)
}

func (m *Mesh) Scale(magnitude mgl32.Vec3) {
	m.scale = m.scale.Add(magnitude)
}

func (m *Mesh) updateModel() {
	m.model = mgl32.Ident4()
	m.model = m.model.Mul4(mgl32.Translate3D(m.position[0], m.position[1], m.position[2]))
	m.model = m.model.Mul4(mgl32.HomogRotate3D(mgl32.DegToRad(m.rotation[0]), mgl32.Vec3{1, 0, 0}))
	m.model = m.model.Mul4(mgl32.HomogRotate3D(mgl32.DegToRad(m.rotation[1]), mgl32.Vec3{0, 1, 0}))
	m.model = m.model.Mul4(mgl32.HomogRotate3D(mgl32.DegToRad(m.rotation[2]), mgl32.Vec3{0, 0, 1}))
	m.model = m.model.Mul4(mgl32.Scale3D(m.scale[0], m.scale[1], m.scale[2]))
}

func (m *Mesh) Render(shader *Shader) {
	m.updateModel()
	shader.SetMat4(m.model, "model", false)
	shader.Bind()
	gl.BindVertexArray(m.vao)
	if m.indexCount > 0 {
		gl.DrawElementsWithOffset(gl.TRIANGLES, 3*int32(m.triangleCount), gl.UNSIGNED_INT, 0)
	} else {
		gl.DrawArrays(gl.TRIANGLES, 0, int32(m.vertexCount))
	}
	shader.Unbind()
}
