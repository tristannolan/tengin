package tengin

import (
	"sync"

	"github.com/gdamore/tcell/v3"
)

// Rewrite this to separate out tcell.EventKey
// Input should be tied more to the engine, with tcell only collecting input
// Also I should really use a rune since everything is single string
type input struct {
	mu      sync.RWMutex
	liveKey *tcell.EventKey
	safeKey *tcell.EventKey
	key     rune
}

func newInput() input {
	return input{
		safeKey: &tcell.EventKey{},
		liveKey: &tcell.EventKey{},
	}
}

func (i *input) Str() string {
	i.mu.RLock()
	defer i.mu.RUnlock()

	return i.safeKey.Str()
}

func (i *input) Key() tcell.Key {
	i.mu.RLock()
	defer i.mu.RUnlock()

	return i.safeKey.Key()
}

func (i *input) listen(e *Engine) {
	go func() {
		for {
			ev := <-e.screen.EventQ()

			switch ev := ev.(type) {
			case *tcell.EventResize:
				e.screen.Sync()
			case *tcell.EventKey:
				e.input.setLiveKey(ev)
			}
		}
	}()
}

func (i *input) poll() {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.safeKey = i.liveKey
}

func (i *input) setLiveKey(key *tcell.EventKey) {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.liveKey = key
}
