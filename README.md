<div align="center">
  <img src="logo.png" alt="sandbox-go" width="512"/>

  [![Go](https://img.shields.io/badge/Go-1.21.1-00ADD8?logo=go&logoColor=white)](https://go.dev/)
  [![Dev Container](https://img.shields.io/badge/Dev%20Container-Ready-blue?logo=docker)](https://containers.dev/)
  [![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

  **🐹 A containerized Go learning environment with 21 progressive tutorials, from hello world to interfaces 🚀**

  [Quick Start](#quick-start) · [Examples](#examples) · [Web Server](#web-server) · [SDL2 Graphics](#sdl2-graphics)
</div>

## Features

- **21 Progressive Tutorials** - Numbered examples (001-021) covering Go fundamentals in logical order
- **Zero Setup** - Docker-based Dev Container handles all dependencies automatically
- **Batteries Included** - Go 1.21.1, SDL2 libraries, and VS Code extensions pre-configured
- **Real-World Examples** - HTTP server with graceful shutdown and SDL2 graphics demo

## Quick Start

```bash
# Clone and open in VS Code
git clone https://github.com/tsilva/sandbox-go.git
code sandbox-go
```

When prompted, click **"Reopen in Container"** (or press F1 → "Dev Containers: Reopen in Container").

Run your first example:

```bash
go run examples/001-hello.go
# Output: hello world
```

## Prerequisites

- [Docker](https://www.docker.com/products/docker-desktop/)
- [VS Code](https://code.visualstudio.com/)
- [Dev Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)

## Examples

The `examples/` directory contains 21 progressive tutorials covering Go language features:

| # | File | Topic |
|---|------|-------|
| 001 | `hello.go` | Hello World |
| 002 | `values.go` | Basic Values |
| 003 | `variables.go` | Variables |
| 004 | `constants.go` | Constants |
| 005 | `for.go` | For Loops |
| 006 | `ifelse.go` | If/Else |
| 007 | `switch.go` | Switch Statements |
| 008 | `arrays.go` | Arrays |
| 009 | `slices.go` | Slices |
| 010 | `maps.go` | Maps |
| 011 | `functions.go` | Functions |
| 012 | `multiple-return-values.go` | Multiple Return Values |
| 013 | `variadic-functions.go` | Variadic Functions |
| 014 | `closures.go` | Closures |
| 015 | `recursion.go` | Recursion |
| 016 | `range-types.go` | Range |
| 017 | `pointers.go` | Pointers |
| 018 | `runes.go` | Strings and Runes |
| 019 | `structs.go` | Structs |
| 020 | `interfaces.go` | Interfaces |
| 021 | `enums.go` | Enums |

Run any example:

```bash
go run examples/007-switch.go
```

## Web Server

A production-ready HTTP server example with graceful shutdown:

```bash
go run examples/webserver.go
# Server is running on http://localhost:8080
```

**Endpoints:**
- `GET /` - Welcome page
- `GET /health` - Health check endpoint

The server handles `SIGINT` gracefully with a 30-second shutdown timeout.

## SDL2 Graphics

An SDL2 graphics example demonstrating window creation and event handling:

```bash
go run examples/sdl.go
```

> **Note:** Requires a graphical environment. Won't work in headless containers.

## Benchmark

Simple performance benchmarking example:

```bash
go run benchmark/benchmark.go
```

## Project Structure

```
.
├── .devcontainer/
│   ├── Dockerfile        # Container with Go 1.21.1 + SDL2
│   └── devcontainer.json # VS Code configuration
├── benchmark/
│   └── benchmark.go      # Performance benchmarking
├── examples/
│   ├── 001-hello.go      # Tutorial examples (001-021)
│   ├── ...
│   ├── webserver.go      # HTTP server example
│   └── sdl.go            # SDL2 graphics example
├── go.mod                # Go module definition
└── README.md
```

## Development Environment

The Dev Container includes:

| Component | Version/Details |
|-----------|-----------------|
| Go | 1.21.1 |
| SDL2 | libsdl2-dev |
| User | devuser (uid/gid 1000) |
| Workspace | /workspace |
| GOPATH | /home/devuser/go |

## License

MIT
