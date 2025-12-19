package tengin

// Events should be emitted when a key is pressed
// I should be able to ask the engine:
// if event.Plant() {}
// Maybe???
type event struct {
	name string
	key  string // criteria for triggering event?
}
