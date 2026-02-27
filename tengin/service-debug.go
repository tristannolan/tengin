package tengin

import "github.com/tristannolan/tengin/tengin/internal/systems"

var debugSystem = systems.NewDebugger()

// =========
//
//	Usage
//
// =========

// Log will output a message to the console.
// Enable debugging and debug console logs in engine config not visible.
func Log(msg string) {
	debugSystem.Debug(msg)
}

// =========
//
//	Types
//
// =========

type DebugService struct {
	system *systems.Debugger
}

// ============
//
//  Internal
//
// ============
