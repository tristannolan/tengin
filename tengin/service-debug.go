package tengin

import "github.com/tristannolan/tengin/tengin/internal/systems"

var debugSystem = systems.NewDebugger()

type DebugProfiler struct {
	internal *systems.Profiler
}

func Debug(msg string)
func Info(msg string)
func Warn(msg string)
func Error(msg string)

func PersistentLog(label, msg string)
func DestroyPersistentLog(label string)

func Profiler(label string) *DebugProfiler

// persistent functions
// profiler functions
// output per level functions

// ============
//
//	Internal
//
// ============
func linkDebug(e *Engine) {
	e.debugSystem = debugSystem
}
