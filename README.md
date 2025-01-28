# sandbox-go

A sandbox environment for Go development using Dev Containers.

## Prerequisites

- [Docker](https://www.docker.com/products/docker-desktop/)
- [VS Code](https://code.visualstudio.com/)
- [Dev Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)

## Getting Started

1. Clone this repository
2. Open in VS Code
3. When prompted, click "Reopen in Container"
   - Alternatively, press F1 and select "Dev Containers: Reopen in Container"

## Development Environment

This sandbox includes:
- Go 1.21.1
- Essential build tools
- Git for version control
- VS Code Go extension

The development container provides an isolated, consistent environment for Go development.

## Project Structure

```
.
├── .devcontainer/
│   ├── Dockerfile        # Container configuration
│   └── devcontainer.json # VS Code configuration
└── README.md
```
