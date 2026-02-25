package systems

type Systems struct {
	Debug  *DebugSystem
	Input  *InputSystem
	Audio  *AudioSystem
	Render *RenderSystem
	Screen *ScreenSystem
}

type (
	DebugSystem  struct{} // file logging, console logging
	InputSystem  struct{} // holds tcell input listener
	AudioSystem  struct{}
	RenderSystem struct{} // holds logic to draw to screen
	ScreenSystem struct{} // holds tcell
)
