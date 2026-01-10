package tengin

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/gdamore/tcell/v3"
	"github.com/tristannolan/tengin/cmd"
)

var (
	consoleMessages         = []string{}
	debugMessages           = []debugMsg{}
	persistentDebugMessages = []debugMsg{}
	debugTimers             = map[int]*debugTimer{}
	longestName             = 0
	longestValue            = 0
	nextDebugTimerId        = 0
)

type debug struct {
	enabled        bool
	canvas         *Canvas
	cmd            *cmd.Controller
	bufferingInput bool
	consoleLines   int
}

type debugMsg struct {
	name  string
	value string
}

type debugTimer struct {
	name             string
	id               int
	maxLogs          int
	logCount         int
	total, lastTotal time.Duration
	start, end       time.Time
}

func newDebug(screenWidth, screenHeight int) *debug {
	consoleLines := 10
	cmdController := cmd.NewController()

	ConsoleLog("setting command")
	err := cmdController.Register(cmd.New("set", func() {
		ConsoleLog("CMD: set")
	}))
	if err != nil {
		ConsoleLog(fmt.Sprintf("%v", err))
	}

	return &debug{
		enabled:        true,
		canvas:         newDebugCanvas(screenWidth, screenHeight, consoleLines),
		cmd:            cmdController,
		bufferingInput: false,
		consoleLines:   consoleLines,
	}
}

func newDebugCanvas(screenWidth, screenHeight, consoleLines int) *Canvas {
	width := screenWidth / 2
	height := consoleLines + 1
	bg := NewColor(10, 10, 10)

	box := Box(width, height, bg)
	box.Transform(0, screenHeight-height)

	return box
}

func (d *debug) handleCommandInput(key Key) {
	if key.Value() == "Empty" {
		return
	}

	if key.Value() == d.cmd.Trigger && !d.bufferingInput {
		d.bufferingInput = true
		return
	}

	if d.bufferingInput {
		switch key.SpecialValue() {
		case KeyEnter:
			d.cmd.Execute()
		case KeyBackspace:
			d.cmd.RemoveFromBuffer(1)
		default:
			d.cmd.AppendToBuffer(key.Value())
		}
	}

	if len(d.cmd.Buffer()) == 0 {
		d.bufferingInput = false
	}
}

func newDebugMsg(name string, value any) debugMsg {
	msg := debugMsg{
		name: name,
	}

	switch v := value.(type) {
	case string:
		msg.value = v
	case int32:
		msg.value = strconv.Itoa(int(v))
	case int:
		msg.value = strconv.Itoa(v)
	case float32:
		msg.value = strconv.FormatFloat(float64(v), 'f', 2, 32)
	case float64:
		msg.value = strconv.FormatFloat(v, 'f', 2, 64)
	case bool:
		msg.value = strconv.FormatBool(v)
	default:
		msg.value = ""
	}

	if len(msg.name) > longestName {
		longestName = len(msg.name)
	}
	if len(msg.value) > longestValue {
		longestValue = len(msg.value)
	}

	return msg
}

func (d *debug) updateCanvas() {
	//if len(consoleMessages) <= 0 {
	//	return
	//}

	c := d.canvas

	// Reset frame
	for y := range c.Tiles {
		for _, tile := range c.Tiles[y] {
			tile.Char = ""
		}
	}

	// Console Messages
	wrapped := make([]string, 0)
	for _, msg := range consoleMessages {
		r := strings.Split(msg, "")

		for i := 0; i < len(r); i += c.Width {
			end := i + c.Width
			if end > len(r) {
				end = len(r)
			}
			final := strings.Join(r[i:end], "")
			wrapped = append(wrapped, final)
		}
	}

	if len(wrapped) > d.consoleLines {
		wrapped = wrapped[len(wrapped)-d.consoleLines:]
	}
	if len(consoleMessages) > d.consoleLines {
		consoleMessages = consoleMessages[len(consoleMessages)-d.consoleLines:]
	}

	// Draw Console
	y := c.Height - 1
	for i := len(wrapped) - 1; i >= 0 && y >= 0; i-- {
		line := strings.Split(wrapped[i], "")

		for x := 0; x < len(line) && x < c.Width; x++ {
			c.Tiles[y-1][x].Char = string(line[x])
		}
		y--
	}

	// Draw Command
	cmdBuffer := d.cmd.Buffer()
	if d.bufferingInput {
		c.Tiles[c.Height-1][0].Char = ":"
		for x := 0; x < c.Width-1; x++ {
			tile := c.Tiles[c.Height-1][x+1]
			if x < len(cmdBuffer) {
				tile.Char = string(cmdBuffer[x])
			} else {
				tile.Char = ""
			}
		}
	}

	c.dirty = true
}

func (d debug) draw(s tcell.Screen) {
	if d.enabled == false {
		return
	}

	d.updateCanvas()

	for i := range nextDebugTimerId {
		if _, ok := debugTimers[i]; !ok {
			continue
		}

		t := debugTimers[i]
		DebugLog(t.name, t.getMsg())
	}

	msgs := slices.Concat(debugMessages, persistentDebugMessages)

	w, h := s.Size()
	x := w - longestName - longestValue - 1
	y := h - len(msgs)

	for i, msg := range msgs {
		whitespace := ""
		for range longestName - len(msg.name) {
			whitespace += " "
		}

		s.PutStr(x-1, y+i, fmt.Sprintf("%s%s %s", msg.name, whitespace, msg.value))
	}

	debugMessages = []debugMsg{}
}

func ConsoleLog(msg string) {
	consoleMessages = append(consoleMessages, msg)
}

func ConsoleLogF(format string, a ...any) {
	consoleMessages = append(consoleMessages, fmt.Sprintf(format, a...))
}

func DebugLog(name string, value any) {
	msg := newDebugMsg(name, value)
	debugMessages = append(debugMessages, msg)
}

func PersistentDebugLog(name string, value any) {
	msg := newDebugMsg(name, value)
	persistentDebugMessages = append(persistentDebugMessages, msg)
}

func NewDebugTimer(name string) *debugTimer {
	dt := &debugTimer{
		id:        nextDebugTimerId,
		name:      name,
		maxLogs:   120,
		logCount:  0,
		total:     0,
		lastTotal: 0,
	}
	debugTimers[dt.id] = dt
	nextDebugTimerId++
	return dt
}

func (dt *debugTimer) Start() {
	dt.start = time.Now()
}

func (dt *debugTimer) End() {
	dt.end = time.Now()

	dt.total += dt.end.Sub(dt.start)

	if dt.logCount <= dt.maxLogs {
		dt.logCount++
		return
	}

	dt.lastTotal = dt.total
	dt.total = 0
	dt.logCount = 0
}

func (t *debugTimer) getMsg() string {
	return fmt.Sprintf("%s", t.lastTotal/time.Duration(t.maxLogs))
}
