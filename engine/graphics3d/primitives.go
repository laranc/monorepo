package graphics3d

import "github.com/go-gl/mathgl/mgl32"

type Primitive interface {
	Set(vertices []Vertex, indices []uint32, vertexCount, indexCount, triangleCount int)
	GetVertices() []Vertex
	GetIndices() []uint32
	GetVertexCount() int
	GetIndexCount() int
	GetTriangleCount() int
}

type Quad struct {
	vertices                               []Vertex
	indices                                []uint32
	vertexCount, indexCount, triangleCount int
}

func NewQuad() *Quad {
	quad := new(Quad)
	vertices := []Vertex{
		{Position: mgl32.Vec3{-0.5, 0.5, 0.0}, Color: mgl32.Vec3{1.0, 1.0, 1.0}, Texcoord: mgl32.Vec2{0.0, 1.0}, Normal: mgl32.Vec3{0.0, 0.0, 1.0}},  // Top Left
		{Position: mgl32.Vec3{-0.5, -0.5, 0.0}, Color: mgl32.Vec3{1.0, 1.0, 1.0}, Texcoord: mgl32.Vec2{0.0, 0.0}, Normal: mgl32.Vec3{0.0, 0.0, 1.0}}, // Bottom Left
		{Position: mgl32.Vec3{0.5, -0.5, 0.0}, Color: mgl32.Vec3{1.0, 1.0, 1.0}, Texcoord: mgl32.Vec2{1.0, 0.0}, Normal: mgl32.Vec3{0.0, 0.0, 1.0}},  // Bottom Right
		{Position: mgl32.Vec3{0.5, 0.5, 0.0}, Color: mgl32.Vec3{1.0, 1.0, 1.0}, Texcoord: mgl32.Vec2{1.0, 1.0}, Normal: mgl32.Vec3{0.0, 0.0, 1.0}},   // Top Right
	}
	indices := []uint32{
		0, 1, 2,
		0, 2, 3,
	}
	quad.Set(vertices, indices, 4, 6, 2)
	return quad
}

func (q *Quad) Set(vertices []Vertex, indices []uint32, vertexCount, indexCount, triangleCount int) {
	q.vertices = vertices
	q.indices = indices
	q.vertexCount = vertexCount
	q.indexCount = indexCount
	q.triangleCount = triangleCount
}

func (q *Quad) GetVertices() []Vertex {
	return q.vertices
}

func (q *Quad) GetIndices() []uint32 {
	return q.indices
}

func (q *Quad) GetVertexCount() int {
	return q.vertexCount
}

func (q *Quad) GetIndexCount() int {
	return q.indexCount
}

func (q *Quad) GetTriangleCount() int {
	return q.triangleCount
}

type Triangle struct {
	vertices                               []Vertex
	indices                                []uint32
	vertexCount, indexCount, triangleCount int
}

func NewTriangle() *Triangle {
	triangle := new(Triangle)
	vertices := []Vertex{
		{Position: mgl32.Vec3{-0.5, 0.5, 0.0}, Color: mgl32.Vec3{1.0, 1.0, 1.0}, Texcoord: mgl32.Vec2{0.0, 1.0}, Normal: mgl32.Vec3{0.0, 0.0, 1.0}},  // Top Left
		{Position: mgl32.Vec3{-0.5, -0.5, 0.0}, Color: mgl32.Vec3{1.0, 1.0, 1.0}, Texcoord: mgl32.Vec2{0.0, 0.0}, Normal: mgl32.Vec3{0.0, 0.0, 1.0}}, // Bottom Left
		{Position: mgl32.Vec3{0.5, -0.5, 0.0}, Color: mgl32.Vec3{1.0, 1.0, 1.0}, Texcoord: mgl32.Vec2{1.0, 0.0}, Normal: mgl32.Vec3{0.0, 0.0, 1.0}},  // Bottom Right
	}
	indices := []uint32{
		0, 1, 2,
	}
	triangle.Set(vertices, indices, 3, 3, 1)
	return triangle
}

func (t *Triangle) Set(vertices []Vertex, indices []uint32, vertexCount, indexCount, triangleCount int) {
	t.vertices = vertices
	t.indices = indices
	t.vertexCount = vertexCount
	t.indexCount = indexCount
	t.triangleCount = triangleCount
}

func (t *Triangle) GetVertices() []Vertex {
	return t.vertices
}

func (t *Triangle) GetIndices() []uint32 {
	return t.indices
}

func (t *Triangle) GetVertexCount() int {
	return t.vertexCount
}

func (t *Triangle) GetIndexCount() int {
	return t.indexCount
}

func (t *Triangle) GetTriangleCount() int {
	return t.triangleCount
}

type Pyramid struct {
	vertices                   []Vertex
	vertexCount, triangleCount int
}

func NewPyramid() *Pyramid {
	pyramid := new(Pyramid)
	vertices := []Vertex{
		//Triangle front
		{Position: mgl32.Vec3{0.0, 0.5, 0.0}, Color: mgl32.Vec3{1.0, 0.0, 0.0}, Texcoord: mgl32.Vec2{0.5, 1.0}, Normal: mgl32.Vec3{0.0, 0.0, 1.0}},
		{Position: mgl32.Vec3{-0.50, -0.5, 0.5}, Color: mgl32.Vec3{0.0, 1.0, 0.0}, Texcoord: mgl32.Vec2{0.0, 0.0}, Normal: mgl32.Vec3{0.0, 0.0, 1.0}},
		{Position: mgl32.Vec3{0.50, -0.5, 0.5}, Color: mgl32.Vec3{0.0, 0.0, 1.0}, Texcoord: mgl32.Vec2{1.0, 0.0}, Normal: mgl32.Vec3{0.0, 0.0, 1.0}},

		//Triangle left
		{Position: mgl32.Vec3{0.0, 0.5, 0.0}, Color: mgl32.Vec3{1.0, 1.0, 0.0}, Texcoord: mgl32.Vec2{0.5, 1.0}, Normal: mgl32.Vec3{-1.0, 0.0, 0.0}},
		{Position: mgl32.Vec3{-0.5, -0.5, -0.5}, Color: mgl32.Vec3{0.0, 0.0, 1.0}, Texcoord: mgl32.Vec2{0.0, 0.0}, Normal: mgl32.Vec3{-1.0, 0.0, 0.0}},
		{Position: mgl32.Vec3{-0.5, -0.5, 0.5}, Color: mgl32.Vec3{0.0, 0.0, 1.0}, Texcoord: mgl32.Vec2{1.0, 0.0}, Normal: mgl32.Vec3{-1.0, 0.0, 0.0}},

		//Triangle back
		{Position: mgl32.Vec3{0.0, 0.5, 0.0}, Color: mgl32.Vec3{1.0, 1.0, 0.0}, Texcoord: mgl32.Vec2{0.5, 1.0}, Normal: mgl32.Vec3{0.0, 0.0, -1.0}},
		{Position: mgl32.Vec3{0.5, -0.5, -0.5}, Color: mgl32.Vec3{0.0, 0.0, 1.0}, Texcoord: mgl32.Vec2{0.0, 0.0}, Normal: mgl32.Vec3{0.0, 0.0, -1.0}},
		{Position: mgl32.Vec3{-0.5, -0.5, -0.5}, Color: mgl32.Vec3{0.0, 0.0, 1.0}, Texcoord: mgl32.Vec2{1.0, 0.0}, Normal: mgl32.Vec3{0.0, 0.0, -1.0}},

		//Triangles right
		{Position: mgl32.Vec3{0.0, 0.5, 0.0}, Color: mgl32.Vec3{1.0, 1.0, 0.0}, Texcoord: mgl32.Vec2{0.5, 1.0}, Normal: mgl32.Vec3{1.0, 0.0, 0.0}},
		{Position: mgl32.Vec3{0.5, -0.5, 0.5}, Color: mgl32.Vec3{0.0, 0.0, 1.0}, Texcoord: mgl32.Vec2{0.0, 0.0}, Normal: mgl32.Vec3{1.0, 0.0, 0.0}},
		{Position: mgl32.Vec3{0.5, -0.5, -0.5}, Color: mgl32.Vec3{0.0, 0.0, 1.0}, Texcoord: mgl32.Vec2{1.0, 0.0}, Normal: mgl32.Vec3{1.0, 0.0, 0.0}},
	}
	pyramid.Set(vertices, nil, 12, 0, 4)
	return pyramid
}

func (p *Pyramid) Set(vertices []Vertex, indices []uint32, vertexCount, indexCount, triangleCount int) {
	p.vertices = vertices
	p.vertexCount = vertexCount
	p.triangleCount = triangleCount
}

func (p *Pyramid) GetVertices() []Vertex {
	return p.vertices
}

func (p *Pyramid) GetIndices() []uint32 {
	return nil
}

func (p *Pyramid) GetVertexCount() int {
	return p.vertexCount
}

func (p *Pyramid) GetIndexCount() int {
	return 0
}

func (p *Pyramid) GetTriangleCount() int {
	return p.triangleCount
}
