# The Engine
The engine provides the tools needed to create an interactive program in the
terminal. It interfaces between the program and the operating environment.

## Structure
- Runtime 
    * Lifecycle <- start, loop, shutdown
- Interface to OS/external tools:
    * Draw  <- convenient functions to draw to the terminal
    * Audio <- control audio events from audio library
- Internal tools
- Debugging

type Engine struct {
    mu sync.RWMutex
    config
}

## Lifecycle
The engine is initialised using the default settings. An optional configuration
struct can be provided for custom settings.

To gain any functionality, a struct must be created to implement the engine
loopable interface. This makes the update and draw functions available.

Calling Run() will then start the program execution. After a final setup phase,
the engine is fully booted and will start looping.

init()
    loadConfig(config)
run()
    lifecycle.running = true
	initSystems()
	initServices()
	initRuntiming()
	initFrameContext()
   	loop()
   	 	lifecycle.check()
   	 	runtime.updateTimers()
		ctx := e.newContext()
   	 	update()
   	 		systems.input.poll()
   	 		updatables.Update(ctx) <- user performs their own updates
   	 	draw()
   	 		drawables.Draw(ctx) <- draw operations received through drawService
   	 		systems.renderer.render()
   	 	runtime.waitForNextTick()
shutdown()

