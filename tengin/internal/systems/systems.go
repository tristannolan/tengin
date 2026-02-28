package systems

// We do not include debug since that operates globally
type Systems struct{}

func NewSystems() *Systems {
	s := &Systems{}

	return s
}
