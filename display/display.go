package display

// Display interface for display struct
type Display interface{}

// display struct
type display struct{}

// newDisplay --> Creates a new display struct
func newDisplay() Display {
	return &display{}
}