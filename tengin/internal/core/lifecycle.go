package core

type Lifecycle struct {
	running           bool
	paused            bool
	shutdownRequested bool
}
