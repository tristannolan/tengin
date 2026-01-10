package tengin

type KeyKind uint8

const (
	KeyKindSpecial KeyKind = iota
	KeyKindMouse
	KeyKindString
)

// List of recognised special keys
type SpecialKey uint16

// CTRL and SHIFT are not included in this list
// Terminals don't report the separate key, they get bundled with another.
// i.e. They require an instance for CTRL-A, CTRL-SHIFT-A, CTRL-B, etc.
// This makes them impractical for games I believe
// I'll add them if I stop justifying not doing so
const (
	// Functional
	KeyEmpty SpecialKey = iota // Might be used as a nil event

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
)

var SpecialKeyNames = map[SpecialKey]string{
	// Functional
	KeyEmpty: "Empty",

	// Keyboard
	KeyEnter:     "Enter",
	KeyTab:       "Tab",
	KeyBacktab:   "Backtab",
	KeyEscape:    "Escape",
	KeyCapsLock:  "CapsLock",
	KeyBackspace: "Backspace",

	// Arrows
	KeyUp:    "Up",
	KeyDown:  "Down",
	KeyRight: "Right",
	KeyLeft:  "Left",

	// Function
	KeyF1:  "F1",
	KeyF2:  "F2",
	KeyF3:  "F3",
	KeyF4:  "F4",
	KeyF5:  "F5",
	KeyF6:  "F6",
	KeyF7:  "F7",
	KeyF8:  "F8",
	KeyF9:  "F9",
	KeyF10: "F10",
	KeyF11: "F11",
	KeyF12: "F12",
}

// Used as the primary vessel for key events
type Key struct {
	kind    KeyKind
	string  string
	special SpecialKey
}

// Returns the string value
// In the case of special keys, the equivilant name will be provided
// See specialKey and specialKeyName for available options
func (k Key) Value() string {
	if k.IsString() {
		return k.string
	}

	return SpecialKeyNames[k.special]
}

// Returns the specialKey code
// Will return KeyEmpty if a string value is set
// Use key.value() to get a string value
func (k Key) SpecialValue() SpecialKey {
	if k.IsString() {
		return KeyEmpty
	}

	return k.special
}

// Create new keys
func NewStringKey(v string) Key {
	return Key{
		kind:   KeyKindString,
		string: v,
	}
}

func NewSpecialKey(k SpecialKey) Key {
	return Key{
		kind:    KeyKindSpecial,
		special: k,
	}
}

func NewEmptyKey() Key {
	return Key{
		kind:    KeyKindSpecial,
		special: KeyEmpty,
	}
}

// Check the key kind value
func (k Key) IsString() bool {
	return k.kind == KeyKindString
}

func (k Key) IsSpecial() bool {
	return k.kind == KeyKindSpecial
}

func (k Key) IsEmpty() bool {
	return k.IsSpecial() && k.special == KeyEmpty
}

// List of recognised mouse keys
type MouseKey uint16

const (
	// Functional
	MouseEmpty MouseKey = iota // Might be used as a nil event

	// Primary
	MouseLeft
	MouseCenter
	MouseRight

	// Secondary
	Mouse1
	Mouse2
	Mouse3
	Mouse4
	Mouse5
	Mouse6
	Mouse7
	Mouse8

	// Wheel
	MouseWheelEmpty
	MouseWheelUp
	MouseWheelDown
	MouseWheelLeft
	MouseWheelRight
)

var MouseKeyNames = map[MouseKey]string{
	// Functional
	MouseEmpty: "MouseEmpty",

	// Primary
	MouseLeft:   "MouseLeft",
	MouseCenter: "MouseCenter",
	MouseRight:  "MouseRight",

	// Wheel
	MouseWheelEmpty: "MouseWheelEmpty",
	MouseWheelUp:    "MouseWheelUp",
	MouseWheelDown:  "MouseWheelDown",
	MouseWheelLeft:  "MouseWheelLeft",
	MouseWheelRight: "MouseWheelRight",
}

type Mouse struct {
	x, y  int
	key   MouseKey
	wheel MouseKey
}

func NewMouse(x, y int, key, wheel MouseKey) Mouse {
	return Mouse{
		x:     x,
		y:     y,
		key:   key,
		wheel: wheel,
	}
}

func NewEmptyMouse() Mouse {
	return Mouse{
		key:   MouseEmpty,
		wheel: MouseWheelEmpty,
	}
}

func (m Mouse) IsEmpty() bool {
	return m.IsKeyEmpty() && m.IsWheelEmpty()
}

func (m Mouse) IsKeyEmpty() bool {
	return m.key == MouseEmpty
}

func (m Mouse) IsWheelEmpty() bool {
	return m.wheel == MouseWheelEmpty
}

func (m Mouse) Position() (int, int) {
	return m.x, m.y
}

func (m Mouse) LastPosition() (int, int) {
	return m.x, m.y
}

func (m Mouse) KeyName() string {
	return MouseKeyNames[m.key]
}

func (m Mouse) Key() MouseKey {
	return m.key
}

func (m Mouse) WheelName() string {
	return MouseKeyNames[m.wheel]
}

func (m Mouse) Wheel() MouseKey {
	return m.wheel
}
