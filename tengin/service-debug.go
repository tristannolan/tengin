package tengin

import "github.com/tristannolan/tengin/tengin/internal/systems"

var debugSystem = systems.NewDebug()

type DebugService struct {
	system *systems.Debugger
}

func Debug(msg string)
func Info(msg string)
func Warn(msg string)
func Error(msg string)

func PersistentLog(label, msg string)
func DestroyPersistentLog(label string)

func Profiler(label string) *systems.Profiler

// persistent functions
// profiler functions
// output per level functions

// ============
//
//  Internal
//
// ============
