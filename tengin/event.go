package tengin

// Identifies if the key is a rune or special character
type keyKind uint8

const (
	keyRune keyKind = iota
	keySpecial
)

// List of recognised special keys
type specialKey uint16

// CTRL and SHIFT are not included in this list
// Terminals don't report the separate key, they get bundled with another.
// i.e. They require an instance for CTRL-A, CTRL-SHIFT-A, CTRL-B, etc.
// This makes them impractical for games I believe
const (
	// Keyboard
	KeyEnter specialKey = iota
	KeyTab
	KeyBacktab
	KeyEscape
	KeyCapsLock
	KeyBackspace // Not working???

	// Arrows
	KeyUp
	KeyDown
	KeyRight
	KeyLeft

	// Function
	KeyF1
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12

	// Mouse
	KeyCenter
)

// Used as the primary vessel for key events
type Key struct {
	kind    keyKind
	rune    rune
	special specialKey
}

func newRuneKey(r rune) Key {
	return Key{
		kind: keyRune,
		rune: r,
	}
}

func newSpecialKey(k specialKey) Key {
	return Key{
		kind:    keySpecial,
		special: k,
	}
}
