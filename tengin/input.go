package tengin

import (
	"sync"

	"github.com/gdamore/tcell/v3"
)

type input struct {
	mu               sync.RWMutex
	liveKey          Key
	safeKey          Key
	isResizingScreen bool
}

func newInput() input {
	return input{
		safeKey:          Key{},
		liveKey:          Key{},
		isResizingScreen: false,
	}
}

func (i *input) Key() Key {
	i.mu.RLock()
	defer i.mu.RUnlock()

	return i.safeKey
}

func (i *input) poll() {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.safeKey = i.liveKey
	i.liveKey = Key{}
}

func (i *input) setRuneKey(r rune) {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.liveKey = newRuneKey(r)
}

func (i *input) setSpecialKey(k specialKey) {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.liveKey = newSpecialKey(k)
}

func (i *input) onScreenResizeStart() {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.isResizingScreen = true
}

func (i *input) onScreenResizeComplete() {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.isResizingScreen = false
}

func (i *input) listen(scr tcell.Screen) {
	go func() {
		for {
			ev := <-scr.EventQ()

			switch ev := ev.(type) {
			case *tcell.EventResize:
				i.onScreenResizeStart()
			case *tcell.EventKey:
				if ev.Str() != "" {
					i.setRuneKey(rune(ev.Str()[0]))
					continue
				}

				switch ev.Key() {
				// Keyboard
				case tcell.KeyEnter:
					i.setSpecialKey(KeyEnter)
				case tcell.KeyTab:
					i.setSpecialKey(KeyTab)
				case tcell.KeyBacktab:
					i.setSpecialKey(KeyBacktab)
				case tcell.KeyCapsLock:
					i.setSpecialKey(KeyCapsLock)
				case tcell.KeyEscape:
					i.setSpecialKey(KeyEscape)
				case tcell.KeyDelete, tcell.KeyBackspace, tcell.KeyBackspace2:
					i.setSpecialKey(KeyBackspace)

				// Arrows
				case tcell.KeyUp:
					i.setSpecialKey(KeyUp)
				case tcell.KeyDown:
					i.setSpecialKey(KeyDown)
				case tcell.KeyRight:
					i.setSpecialKey(KeyRight)
				case tcell.KeyLeft:
					i.setSpecialKey(KeyLeft)

				// Function
				case tcell.KeyF1:
					i.setSpecialKey(KeyF1)
				case tcell.KeyF2:
					i.setSpecialKey(KeyF2)
				case tcell.KeyF3:
					i.setSpecialKey(KeyF3)
				case tcell.KeyF4:
					i.setSpecialKey(KeyF4)
				case tcell.KeyF5:
					i.setSpecialKey(KeyF5)
				case tcell.KeyF6:
					i.setSpecialKey(KeyF6)
				case tcell.KeyF7:
					i.setSpecialKey(KeyF7)
				case tcell.KeyF8:
					i.setSpecialKey(KeyF8)
				case tcell.KeyF9:
					i.setSpecialKey(KeyF9)
				case tcell.KeyF10:
					i.setSpecialKey(KeyF10)
				case tcell.KeyF11:
					i.setSpecialKey(KeyF11)
				case tcell.KeyF12:
					i.setSpecialKey(KeyF12)

				// Mouse
				case tcell.KeyCenter:
					i.setSpecialKey(MouseCenter)
				}
			}
		}
	}()
}
