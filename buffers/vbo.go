package buffers

// VBO inteface for vao struct
type VBO interface{}

// vao struct
type vbo struct{}

// NewVBO --> Creates a new vbo struct
func NewVBO() VAO {
	return &vbo{}
}
