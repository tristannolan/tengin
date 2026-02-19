# Concepts
## Purpose of this document
Providing meaning to key interfaces within the engine.

What should be covered:
- What it represents / reason for existing
- What it owns
- What it does not own
- Key properties and constraints
- Relationship to other concepts

## Table of contents
- Scene
- Node
- Canvas
- Control
- Transform
- Render pipeline
- Input system

## Engine
### Purpose
The Engine is the runtime orchestrator of the system.

It does not perform domain work itself. Instead, it coordinates and schedules
specialised subsystems so they operate together in a consistent lifecycle.

It is the primary entry point exposed to the user and represents the top-level
runtime boundary of the engine.

### Owns
The Engine owns responsibilities that are global, or life-cycle related.

#### Lifecycle Management
- Clean startup and shutdown
- Runtime state transitions (running, paused, stopped)

#### Main loop control
- Frame timing
- Runtime execution order (poll input, update, draw, render, debug...)

#### Global configuration
- Configuration struct
- Runtime-adjustable settings

#### Platform interfaces
- Terminal drawing
- Input collection

#### Subsystem coordination
- Rendering system
- Input system
- Scene management
- Debug systems

#### Debug system
- Profiling tools
- Logging pipelines (file, console)
- Command console (for dev level engine management)
- Debug overlays (console)

### Does not own
The Engine does not own domain or subsystem logic.

### Relationships
The Engine is at the top of the hierarchy. 
It is parent and orchestrator of all.

### Lifecycle
#### Phases
- Configuration
- Initalisation
- Run
- Pause (optional)
- Shutdown

#### Configuration
- Load default configuration
- Load user configuration
- Apply settings

#### Initalisation
- Initalise platform interfaces
- Create subsystems
- Prepare runtime state

#### Run
- Execute main loop

#### Pause
- Suspend updates
- Maintain runtime state

#### Shutdown
- Stop subsystems
- Release resource
- Restore platform state
- Output debug/panic messages

### Questions

## Scene
### Purpose
### Owns
### Does not own
### Relationships
### Lifecycle
### Questions

## Node
### Purpose
A Node is an interfacing surface to render and control objects on screen.

Nodes are bound to a specific scene, and can contain the follow properties:
- Canvas
- Control
- Layout
- Update()
- Draw()

### Owns
### Does not own
### Relationships
### Lifecycle
### Questions
