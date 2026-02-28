package systems

import "time"

func NewDebugger() *Debugger {
	d := &Debugger{
		enabled: true,
		output: map[logLevel][]logOutput{
			logLevelDebug: {logOutputConsole},
			logLevelInfo:  {logOutputConsole},
			logLevelWarn:  {logOutputConsole, logOutputFile},
			logLevelError: {logOutputConsole, logOutputFile},
		},
	}
	return d
}

func NewProfiler() *Profiler {
	return &Profiler{}
}

type Debugger struct {
	enabled bool

	bufEntries        []debugEntry
	persistentEntries map[string]debugEntry
	profilers         map[string]*Profiler

	output map[logLevel][]logOutput
}

type Profiler struct {
	entry      profileEntry
	total      time.Duration
	count      int
	maxEntries int
}

type debugEntry struct {
	lvl logLevel
	msg string
}

type profileEntry struct {
	label      string
	running    bool
	start, end time.Time
}

type logLevel int

const (
	logLevelDebug logLevel = iota
	logLevelInfo
	logLevelWarn
	logLevelError
)

type logOutput int

const (
	logOutputConsole logOutput = iota
	logOutputFile
)

// =========
//
//	Usage
//
// =========

// BeginFrame will clear the entries buffer and prepare for new logs.
// Persistent entries and profilers will not be cleared.
func (d *Debugger) BeginFrame() {
}

// EndFrame will process the buffer and return entries.
func (d *Debugger) EndFrame() []debugEntry {
	return []debugEntry{}
}

// use log to output with the correct level
func (d *Debugger) Debug(msg string) {
	d.log(logLevelDebug, msg)
}

func (d *Debugger) Info(msg string) {
	d.log(logLevelInfo, msg)
}

func (d *Debugger) Warn(msg string) {
	d.log(logLevelWarn, msg)
}

func (d *Debugger) Error(msg string) {
	d.log(logLevelError, msg)
}

// Logs with name value pair on screen.
// Intended for long running, status type data.
func (d *Debugger) PersistentLog(label, msg string)   {}
func (d *Debugger) DestroyPersistentLog(label string) {}

// Profiler is used to time function execution.
// It is recommended store these on a persistent data structure so they are not
// being created each tick.
// Profilers are automatically added to the debugger.
func (d *Debugger) Profiler(label string) *Profiler {
	return &Profiler{}
}

func (p *Profiler) Start() {}

func (p *Profiler) End() profileEntry {
	return profileEntry{}
}

func (p *Profiler) Average() {
}

// ============
//
//	Internal
//
// ============

func (d *Debugger) log(lvl logLevel, msg string) {
	if !d.enabled {
		return
	}

	e := debugEntry{
		lvl: lvl,
		msg: msg,
	}

	d.bufEntries = append(d.bufEntries, e)
}
