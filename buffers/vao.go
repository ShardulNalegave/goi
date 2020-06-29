package buffers

import "github.com/go-gl/gl/v4.1-core/gl"

// VAO inteface for vao struct
type VAO interface {
	init()
	GetID() uint32
	Use(func())
}

// vao struct
type vao struct {
	id uint32
}

func (v *vao) init() {
	gl.GenVertexArrays(0, &v.id)
}

func (v *vao) GetID() uint32 {
	return v.id
}

func (v *vao) Use(handler func()) {
	gl.BindVertexArray(v.id)
	handler()
	gl.BindVertexArray(0)
}

// NewVAO --> Creates a new vao struct
func NewVAO() VAO {
	v := vao{}
	v.init()
	return &v
}
