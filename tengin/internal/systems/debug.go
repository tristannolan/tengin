package systems

var messages = []debugMsg{}

// =========
//
//	Usage
//
// =========

// One shot log to the console.
//
// Usage: d.Debug("my console message")
func (d *Debugger) Debug(msg debugMsg) {
	d.log(logLevelDebug, msg)
}

func (d *Debugger) Info(msg debugMsg) {
	d.log(logLevelInfo, msg)
}

func (d *Debugger) Warn(msg debugMsg) {
	d.log(logLevelWarn, msg)
}

func (d *Debugger) Error(msg debugMsg) {
	d.log(logLevelError, msg)
}

// Logs with name value pair on screen.
// Intended for long running, status type data.
// This function can be called every tick and the value will be updated instead
// of creating new logs.
// See d.DestroyPersistentLog() to remove them.
//
// PersistentLog("tps", e.runtime.tps)
func (d *Debugger) PersistentLog(label, msg string) {
}

func (d *Debugger) DestroyPersistentLog(label string) {
}

// One shot log to a file.
// See Debug config to set log location.
// Usage: LogFile("this goes to a file")
func (d *Debugger) DebugFile(msg string) {
}

// Profile is used to time function execution.
// It is recommended store these on a persistent data structure so they are not
// being created each tick.
//
// Usage: p := Profile("unique label")
// Usage: p.Start()
// Usage: p.End()
func (d *Debugger) Profile(msg string) {
}

// =========
//
//	Types
//
// =========
type Debugger struct {
	enabled bool
	output  map[logLevel]debugOutputs
}

type debugMsg struct{}

type logLevel int

const (
	logLevelDebug logLevel = iota
	logLevelInfo
	logLevelWarn
	logLevelError
	logLevelConsole
	logLevelFile
)

// new type instance
func NewDebugger() *Debugger {
	d := &Debugger{
		enabled: true,
	}
	return d
}

// ============
//
//	Internal
//
// ============

func (d *Debugger) log(lvl logLevel, msg debugMsg) {
	switch lvl {
	case logLevelDebug:
	case logLevelInfo:
	case logLevelError:
	}
}
