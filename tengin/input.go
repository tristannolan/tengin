package tengin

import (
	"sync"

	"github.com/gdamore/tcell/v3"
)

type input struct {
	mu      sync.RWMutex
	liveKey *tcell.EventKey
	safeKey *tcell.EventKey
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

func (i *input) Raw() *tcell.EventKey {
	i.mu.RLock()
	defer i.mu.RUnlock()

	return i.safeKey
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
