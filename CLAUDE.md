# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go learning sandbox environment configured with Dev Containers. It contains numbered example files (001-021) demonstrating Go language features, plus specialized examples for web servers and SDL2 graphics. The project uses Go 1.21.1 and includes the SDL2 library for graphics programming.

## Development Environment

The project runs in a Docker-based Dev Container with:
- Go 1.21.1
- SDL2 development libraries (libsdl2-dev)
- Non-root user "devuser" (uid/gid 1000)
- Workspace at `/workspace`
- GOPATH at `/home/devuser/go`

## Running Code

Run individual example files:
```bash
go run examples/001-hello.go
go run examples/webserver.go
```

Run the benchmark:
```bash
go run benchmark/benchmark.go
```

For the SDL2 example (requires graphical environment):
```bash
go run examples/sdl.go
```

Install dependencies:
```bash
go mod download
```

## Project Structure

- `examples/` - Numbered tutorial files (001-021) covering Go basics from hello world through interfaces, plus specialized examples:
  - `001-hello.go` through `021-enums.go` - Sequential Go language feature demonstrations
  - `webserver.go` - HTTP server with graceful shutdown and health check endpoint
  - `sdl.go` - SDL2 graphics window with event handling
- `benchmark/` - Simple performance benchmarking example
- `.devcontainer/` - Docker configuration for dev environment

## Important Notes

- README.md must be kept up to date with any significant project changes
- The SDL2 example requires a graphical environment and won't work in headless containers
- The web server listens on port 8080 and includes a `/health` endpoint
- No test files exist in this repository currently
