package systems

import "github.com/gdamore/tcell/v3"

type TcellTermDriver struct {
	scr tcell.Screen
}

func NewTcellTermDriver() (*TcellTermDriver, error) {
	scr, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}

	err = scr.Init()
	if err != nil {
		return nil, err
	}

	t := &TcellTermDriver{
		scr: scr,
	}

	return t, nil
}

// Interface: tengin.SystemLifecycle
func (t *TcellTermDriver) Init() error {
	t.scr.EnableMouse()
	t.scr.EnableFocus()

	return nil
}

func (t *TcellTermDriver) Stop() {
	t.scr.Fini()
}

// Interface: tengin.Renderer
func (t TcellTermDriver) Show() {
	// t.scr.PutStr(0, 0, "test")
	// t.scr.Show()
}

func (t TcellTermDriver) SetTile(x, y int, ch string) {}

func (t TcellTermDriver) Size() (w, h int) {
	return 0, 0
}

// Interface: tengin.Input
func (t TcellTermDriver) Poll() {}

func (t TcellTermDriver) Listen() {
	// go func() {
	//	for {
	//		ev := <-t.scr.EventQ()

	//		switch ev := ev.(type) {
	//		case *tcell.EventKey:
	//			switch ev.Key() {
	//			// Keyboard
	//			case tcell.KeyEscape:
	//				t.scr.Fini()
	//			}
	//		}
	//	}
	//}()
}
