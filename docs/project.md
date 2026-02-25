# Project
## File Structure

docs/
demos/
    01-basic-instance/
    02-configure/
    03-debug/
    04-input/
    05-draw/
    06-control/
tengin/
    tengin.go               <-  Public facing api
    config.go
    context.go
    service-debug.go
    service-input.go
    service-render.go
    internal/               <-  Not user accessible
        systems/
            debug.go
            input.go
            render.go
            screen.go
        runtime/
        lifecycle/

