package buffers

// VBO inteface for vao struct
type VBO interface {
	GetID() uint32
}

// vao struct
type vbo struct {
	id uint32
}

func (v *vbo) GetID() uint32 {
	return v.id
}

// NewVBO --> Creates a new vbo struct
func NewVBO() VBO {
	return &vbo{}
}
