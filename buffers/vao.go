package buffers

// VAO inteface for vao struct
type VAO interface{}

// vao struct
type vao struct{}

// NewVAO --> Creates a new vao struct
func NewVAO() VAO {
	return &vao{}
}
