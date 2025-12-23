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
	// Functional
	KeyEmpty specialKey = iota // Used as a nil key event

	// Keyboard
	KeyEnter
	KeyTab
	KeyBacktab // SHIFT + TAB
	KeyEscape
	KeyCapsLock
	KeyBackspace

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
	MouseCenter
)

// Used as the primary vessel for key events
type Key struct {
	kind    keyKind
	rune    rune
	special specialKey
}

// Create new keys
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

func newEmptyKey() Key {
	return Key{
		kind:    keySpecial,
		special: KeyEmpty,
	}
}

// Check the key kind value
func (k Key) isRuneKey() bool {
	return k.kind == keyRune
}

func (k Key) isSpecialKey() bool {
	return k.kind == keySpecial
}

func (k Key) getRuneValue() rune {
	if k.isRuneKey() {
		return k.rune
	}

	return 0
}

func (k Key) getSpecialValue() specialKey {
	if k.isSpecialKey() {
		return k.special
	}

	return KeyEmpty
}
