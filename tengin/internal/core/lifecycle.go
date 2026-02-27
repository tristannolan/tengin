package core

type Lifecycle struct {
	running bool
	paused  bool
}

func NewLifecycle() *Lifecycle {
	l := &Lifecycle{
		running: false,
		paused:  false,
	}

	return l
}

func (l *Lifecycle) Run()
func (l *Lifecycle) Running() bool

func (l *Lifecycle) Pause()
func (l *Lifecycle) Unpause()
func (l *Lifecycle) Paused() bool

func (l *Lifecycle) RequestShutdown()
func (l *Lifecycle) Shutdown() // call any methods right before closing

func (l *Lifecycle) ShouldUpdate() bool
func (l *Lifecycle) ShouldDraw() bool
