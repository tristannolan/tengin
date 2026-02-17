package tengin

import (
	"cmp"
	"slices"
)

type ControlManager struct {
	Controls []*Control
	dirty    bool
}

type Control struct {
	width, height int
	x, y, z       int
	transform     *Transform
	manager       *ControlManager
	dirty         bool
	hover         bool
	Click         func()
	Hover         func()
	HoverOff      func()
	Key           func(key Key)
}

func newControlManager() *ControlManager {
	return &ControlManager{
		Controls: []*Control{},
		dirty:    true,
	}
}

func NewControl(width, height int) *Control {
	c := &Control{
		width:     width,
		height:    height,
		x:         0,
		y:         0,
		z:         0,
		transform: NewTransform(0, 0),
		manager:   nil,
		dirty:     true,
		hover:     false,
		Click:     func() {},
		Hover:     func() {},
		HoverOff:  func() {},
		Key:       func(key Key) {},
	}

	return c
}

// ===================
//
//	Control Manager
//
// ===================
func (cm *ControlManager) Add(c ...*Control) {
	for _, ctrl := range c {
		ctrl.assignManager(cm)
		cm.Controls = append(cm.Controls, ctrl)
	}
}

func (cm *ControlManager) Remove(c ...*Control) {
	if len(cm.Controls) == 0 {
		return
	}

	toRemove := make(map[*Control]struct{}, len(c))
	toRemain := make([]*Control, len(cm.Controls)-len(c))

	for _, control := range c {
		toRemove[control] = struct{}{}
	}

	for _, control := range cm.Controls {
		if _, found := toRemove[control]; found {
			continue
		}
		toRemain = append(toRemain, control)
	}
}

func (cm *ControlManager) HitKeys(key Key) {
	for _, ctrl := range cm.Controls {
		ctrl.Key(key)
	}
}

func (cm ControlManager) Dirty() bool {
	return cm.dirty
}

func (cm *ControlManager) MarkDirty() {
	if cm.dirty == true {
		return
	}
	cm.dirty = true
}

func (cm *ControlManager) MarkClean() {
	cm.dirty = false
	for _, ctrl := range cm.Controls {
		ctrl.dirty = false
	}
}

func (cm *ControlManager) Sort() {
	slices.SortStableFunc(cm.Controls, func(a, b *Control) int {
		return cmp.Compare(a.z, b.z)
	})
	cm.MarkClean()
}

// ===========
//
//	Control
//
// ===========
func (c Control) Z() int {
	return c.z
}

func (c *Control) SetZ(z int) {
	c.z = z
	c.MarkDirty()
}

func (c *Control) Dirty() bool {
	return c.dirty
}

func (c *Control) ContainsPoint(x, y int) bool {
	return c.bounds().Contains(x, y)
}

func (c *Control) SetClickAction(f func()) {
	c.Click = f
}

func (c *Control) SetHoverAction(f func()) {
	c.Hover = f
}

func (c *Control) SetHoverOffAction(f func()) {
	c.HoverOff = f
}

func (c *Control) SetKeyAction(f func(key Key)) {
	c.Key = f
}

// A canvas will use a locally bound transform unless otherwise specified.
// Assign a new one if the transform must be shared elsewhere.
func (c *Control) AssignTransform(t *Transform) {
	c.transform = t
	c.MarkDirty()
}

func (c *Control) GetTransform() (int, int) {
	return c.transform.x, c.transform.y
}

func (c *Control) bounds() Rect {
	minX := c.x + c.transform.x
	minY := c.y + c.transform.y
	maxX := minX + c.width - 1
	maxY := minY + c.height - 1
	return NewRect(minX, minY, maxX, maxY)
}

func (c *Control) assignManager(cm *ControlManager) {
	c.manager = cm
}

func (c *Control) MarkDirty() {
	if c.dirty == true {
		return
	}

	c.dirty = true
	c.manager.MarkDirty()
}
