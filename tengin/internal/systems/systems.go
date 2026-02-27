package systems

// We do not include debug since that operates globally
type Systems struct {
	Input  *Input
	Audio  *Audio
	Render *Render
	Screen *Screen
}
