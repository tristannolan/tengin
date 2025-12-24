package tengin

import (
	"sync"

	"github.com/gdamore/tcell/v3"
)

type liveInput struct {
	mu sync.RWMutex

	key        Key
	mouseKey   Mouse
	mouseWheel Mouse

	isScreenResizing bool
	isScreenFocused  bool
}

func newLiveInput() liveInput {
	return liveInput{
		key:        NewEmptyKey(),
		mouseKey:   NewEmptyMouse(),
		mouseWheel: NewEmptyMouse(),

		isScreenResizing: false,
		isScreenFocused:  true,
	}
}

type input struct {
	key     Key
	lastKey Key

	mouseKey     Mouse
	lastMouseKey Mouse

	mouseWheel     Mouse
	lastMouseWheel Mouse

	isScreenResizing bool
	isScreenFocused  bool
}

func newInput() input {
	return input{
		key:     NewEmptyKey(),
		lastKey: NewEmptyKey(),

		mouseKey:     NewEmptyMouse(),
		lastMouseKey: NewEmptyMouse(),

		mouseWheel:     NewEmptyMouse(),
		lastMouseWheel: NewEmptyMouse(),

		isScreenResizing: false,
		isScreenFocused:  true,
	}
}

func (i *input) poll(live *liveInput) {
	// Key input
	i.key = live.key
	if !i.key.IsEmpty() {
		i.lastKey = i.key
	}
	live.key = NewEmptyKey()

	// Mouse input
	i.mouseKey = live.mouseKey
	if !i.mouseKey.IsKeyEmpty() {
		i.lastMouseKey = i.mouseKey
	}
	live.mouseKey = NewEmptyMouse()

	i.mouseWheel = live.mouseWheel
	if !i.mouseWheel.IsWheelEmpty() {
		i.lastMouseWheel = i.mouseWheel
	}
	live.mouseWheel = NewEmptyMouse()
}

func (live *liveInput) setStringKey(v string) {
	live.mu.Lock()
	defer live.mu.Unlock()

	live.key = NewStringKey(v)
}

func (live *liveInput) setSpecialKey(k SpecialKey) {
	live.mu.Lock()
	defer live.mu.Unlock()

	live.key = NewSpecialKey(k)
}

func (live *liveInput) setMouse(x, y int, key, wheel MouseKey) {
	live.mu.Lock()
	defer live.mu.Unlock()

	live.mouseKey = NewMouse(x, y, key, wheel)
}

func (live *liveInput) onScreenResizeStart() {
	live.mu.Lock()
	defer live.mu.Unlock()

	live.isScreenResizing = true
}

func (live *liveInput) onScreenResizeComplete() {
	live.mu.Lock()
	defer live.mu.Unlock()

	live.isScreenResizing = false
}

func (live *liveInput) setScreenFocus(b bool) {
	live.mu.Lock()
	defer live.mu.Unlock()

	live.isScreenFocused = b
}

func (live *liveInput) listen(scr tcell.Screen) {
	go func() {
		for {
			ev := <-scr.EventQ()

			switch ev := ev.(type) {
			case *tcell.EventResize:
				live.onScreenResizeStart()
			case *tcell.EventFocus:
				live.setScreenFocus(ev.Focused)
			case *tcell.EventKey:
				if ev.Str() != "" {
					live.setStringKey(ev.Str())
					continue
				}

				switch ev.Key() {
				// Keyboard
				case tcell.KeyEnter:
					live.setSpecialKey(KeyEnter)
				case tcell.KeyTab:
					live.setSpecialKey(KeyTab)
				case tcell.KeyBacktab:
					live.setSpecialKey(KeyBacktab)
				case tcell.KeyCapsLock:
					live.setSpecialKey(KeyCapsLock)
				case tcell.KeyEscape:
					live.setSpecialKey(KeyEscape)
				case tcell.KeyDelete, tcell.KeyBackspace, tcell.KeyBackspace2:
					live.setSpecialKey(KeyBackspace)

				// Arrows
				case tcell.KeyUp:
					live.setSpecialKey(KeyUp)
				case tcell.KeyDown:
					live.setSpecialKey(KeyDown)
				case tcell.KeyRight:
					live.setSpecialKey(KeyRight)
				case tcell.KeyLeft:
					live.setSpecialKey(KeyLeft)

				// Function
				case tcell.KeyF1:
					live.setSpecialKey(KeyF1)
				case tcell.KeyF2:
					live.setSpecialKey(KeyF2)
				case tcell.KeyF3:
					live.setSpecialKey(KeyF3)
				case tcell.KeyF4:
					live.setSpecialKey(KeyF4)
				case tcell.KeyF5:
					live.setSpecialKey(KeyF5)
				case tcell.KeyF6:
					live.setSpecialKey(KeyF6)
				case tcell.KeyF7:
					live.setSpecialKey(KeyF7)
				case tcell.KeyF8:
					live.setSpecialKey(KeyF8)
				case tcell.KeyF9:
					live.setSpecialKey(KeyF9)
				case tcell.KeyF10:
					live.setSpecialKey(KeyF10)
				case tcell.KeyF11:
					live.setSpecialKey(KeyF11)
				case tcell.KeyF12:
					live.setSpecialKey(KeyF12)

				// Mouse
				// case tcell.KeyCenter:
				// 	live.setSpecialKey(MouseCenter)

				// Default to empty key
				default:
					live.setSpecialKey(KeyEmpty)
				}
			case *tcell.EventMouse:
				x, y := ev.Position()
				button := ev.Buttons()

				var key MouseKey
				var wheel MouseKey

				switch {
				case button&tcell.WheelUp != 0:
					wheel = MouseWheelUp
				case button&tcell.WheelDown != 0:
					wheel = MouseWheelDown
				case button&tcell.WheelLeft != 0:
					wheel = MouseWheelLeft
				case button&tcell.WheelRight != 0:
					wheel = MouseWheelRight
				}

				switch ev.Buttons() {
				case tcell.Button1:
					key = MouseLeft
				case tcell.Button2:
					key = MouseRight
				case tcell.Button3:
					key = MouseCenter
				}

				live.setMouse(x, y, key, wheel)
			}
		}
	}()
}
