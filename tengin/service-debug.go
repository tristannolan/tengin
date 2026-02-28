package tengin

import "github.com/tristannolan/tengin/tengin/internal/systems"

var debugSystem = systems.NewDebugger()

type DebugProfiler struct {
	internal *systems.Profiler
}

func Debug(msg string) {}
func Info(msg string)  {}
func Warn(msg string)  {}
func Error(msg string) {}

func Profiler(label string) *DebugProfiler {
	return &DebugProfiler{
		internal: systems.NewProfiler(),
	}
}

// ============
//
//	Internal
//
// ============

func linkDebug(e *Engine) {
	e.debugSystem = debugSystem
}
