# Go Installation Guide

This guide covers installing Go on different operating systems and setting up your development environment.

## Table of Contents

- [System Requirements](#system-requirements)
- [Installation Methods](#installation-methods)
- [Environment Setup](#environment-setup)
- [Verification](#verification)
- [Troubleshooting](#troubleshooting)

## System Requirements

- **Operating System**: Linux, macOS, Windows, FreeBSD
- **Architecture**: amd64, 386, arm64, arm
- **Memory**: At least 1GB RAM
- **Disk Space**: 500MB for Go installation

## Installation Methods

### Method 1: Official Binary (Recommended)

#### Linux/macOS

1. Download the latest version from [golang.org](https://golang.org/dl/)
2. Extract and install:

```bash
# Download (replace with latest version)
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz

# Remove old installation and extract
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
```

#### Windows

1. Download the MSI installer from [golang.org](https://golang.org/dl/)
2. Run the installer and follow the prompts
3. Go will be installed to `C:\Program Files\Go`

### Method 2: Package Manager

#### Fedora/RHEL/CentOS

```bash
sudo dnf install golang
```

#### Ubuntu/Debian

```bash
sudo apt update
sudo apt install golang-go
```

#### macOS (Homebrew)

```bash
brew install go
```

#### Arch Linux

```bash
sudo pacman -S go
```

### Method 3: Version Manager (gvm)

For managing multiple Go versions:

```bash
# Install gvm
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)

# Install and use Go version
gvm install go1.21.5
gvm use go1.21.5 --default
```

## Environment Setup

### Setting PATH and Environment Variables

#### Bash/Zsh (Linux/macOS)

Add to your `~/.bashrc`, `~/.zshrc`, or `~/.profile`:

```bash
# Go installation path
export PATH=$PATH:/usr/local/go/bin

# Go workspace (optional with modules)
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

Apply changes:

```bash
source ~/.bashrc  # or ~/.zshrc
```

#### Windows

1. Open System Properties → Advanced → Environment Variables
2. Add to System PATH: `C:\Program Files\Go\bin`
3. Create GOPATH variable: `C:\Users\YourName\go`

#### Fish Shell

Add to `~/.config/fish/config.fish`:

```fish
set -gx PATH $PATH /usr/local/go/bin
set -gx GOPATH $HOME/go
set -gx GOBIN $GOPATH/bin
set -gx PATH $PATH $GOBIN
```

### Workspace Setup

Create your Go workspace directories:

```bash
mkdir -p $HOME/go/{bin,src,pkg}
```

**Note**: With Go modules (Go 1.11+), you can work outside of GOPATH, but it's still useful for tools and legacy projects.

## Verification

### Check Installation

```bash
go version
```

Expected output:

```
go version go1.21.5 linux/amd64
```

### Check Environment

```bash
go env
```

Key variables to verify:

- `GOROOT`: Go installation directory
- `GOPATH`: Your workspace directory
- `GOOS`: Your operating system
- `GOARCH`: Your architecture

### Test Installation

Create a test program:

```bash
mkdir hello-world && cd hello-world
go mod init hello-world
```

Create `main.go`:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Run it:

```bash
go run main.go
```

## Development Tools (Optional)

Install useful development tools:

```bash
# Language server for editor support
go install golang.org/x/tools/gopls@latest

# Debugger
go install github.com/go-delve/delve/cmd/dlv@latest

# Import formatter
go install golang.org/x/tools/cmd/goimports@latest

# Static analysis
go install honnef.co/go/tools/cmd/staticcheck@latest
```

## Editor Setup

### VS Code

1. Install the "Go" extension by Google
2. Open Command Palette (Ctrl+Shift+P)
3. Run "Go: Install/Update Tools"

### Vim/Neovim

Install vim-go plugin:

```vim
Plug 'fatih/vim-go'
```

### GoLand

JetBrains GoLand comes with built-in Go support.

## Troubleshooting

### Common Issues

#### "go: command not found"

- Check if Go is in your PATH
- Restart your terminal
- Verify installation location

#### Permission denied errors

```bash
# Fix permissions for Go workspace
chmod -R 755 $GOPATH
```

#### GOPATH issues

- With Go modules, GOPATH is less important
- Ensure GOPATH is set if working with legacy code
- Don't put module-based projects in GOPATH/src

#### Proxy issues

If behind a corporate firewall:

```bash
go env -w GOPROXY=direct
go env -w GOSUMDB=off
```

### Getting Help

- Official documentation: https://golang.org/doc/
- Go community: https://gophers.slack.com/
- Stack Overflow: Tag your questions with `go`

## Next Steps

After installation:

1. Read the [Go Commands Guide](GO_COMMANDS.md)
2. Complete the [Tour of Go](https://tour.golang.org/)
3. Try the [Go by Example](https://gobyexample.com/) tutorials
4. Explore the [standard library](https://pkg.go.dev/std)

## Updating Go

### Binary Installation

1. Download new version
2. Remove old installation: `sudo rm -rf /usr/local/go`
3. Extract new version: `sudo tar -C /usr/local -xzf go1.x.x.tar.gz`

### Package Manager

```bash
# Fedora
sudo dnf update golang

# Ubuntu
sudo apt update && sudo apt upgrade golang-go

# macOS
brew upgrade go
```

Remember to restart your terminal or source your shell configuration after updating!

````

```markdown:docs/GO_COMMANDS.md
# Go Commands Guide for Beginners

This guide covers essential Go commands that every beginner should know, organized by use case with practical examples.

## Table of Contents
- [Basic Commands](#basic-commands)
- [Module Management](#module-management)
- [Building and Running](#building-and-running)
- [Testing](#testing)
- [Code Quality](#code-quality)
- [Package Management](#package-management)
- [Debugging and Profiling](#debugging-and-profiling)
- [Environment and Configuration](#environment-and-configuration)

## Basic Commands

### `go version`
Check your Go installation version.

```bash
go version
````

**Output**: `go version go1.21.5 linux/amd64`

**When to use**: Verify installation, check compatibility with tutorials/libraries.

### `go help`

Get help for any Go command.

```bash
go help
go help build
go help mod
```

**When to use**: When you need to understand command options and usage.

## Module Management

### `go mod init`

Initialize a new Go module (project).

```bash
go mod init myproject
go mod init github.com/username/myproject
```

**Creates**: `go.mod` file that tracks dependencies.

**When to use**: Starting any new Go project.

**Example**:

```bash
mkdir calculator && cd calculator
go mod init calculator
```

### `go mod tidy`

Clean up your module dependencies.

```bash
go mod tidy
```

**What it does**:

- Removes unused dependencies
- Adds missing dependencies
- Updates `go.mod` and `go.sum`

**When to use**: After adding/removing imports, before committing code.

### `go mod download`

Download dependencies without building.

```bash
go mod download
go mod download github.com/gorilla/mux
```

**When to use**: Pre-download dependencies, CI/CD pipelines.

### `go get`

Add, upgrade, or remove dependencies.

```bash
# Add a dependency
go get github.com/gorilla/mux

# Get specific version
go get github.com/gorilla/mux@v1.8.0

# Upgrade to latest
go get -u github.com/gorilla/mux

# Remove dependency
go get github.com/gorilla/mux@none
```

**When to use**: Adding external libraries to your project.

## Building and Running

### `go run`

Compile and run Go programs directly.

```bash
go run main.go
go run .
go run *.go
```

**When to use**:

- Development and testing
- Quick scripts
- Learning Go

**Example**:

```bash
echo 'package main; import "fmt"; func main() { fmt.Println("Hello!") }' > hello.go
go run hello.go
```

### `go build`

Compile Go programs into executable binaries.

```bash
# Build current directory
go build

# Build specific file
go build main.go

# Build with custom name
go build -o myapp

# Build for different OS/architecture
GOOS=windows GOARCH=amd64 go build -o myapp.exe
```

**When to use**: Creating distributable executables.

**Example**:

```bash
go build -o calculator
./calculator
```

### `go install`

Compile and install packages/commands.

```bash
# Install current project to $GOBIN
go install

# Install external tool
go install golang.org/x/tools/cmd/goimports@latest
```

**When to use**: Installing command-line tools, deploying applications.

## Testing

### `go test`

Run tests in your project.

```bash
# Test current package
go test

# Test all packages
go test ./...

# Verbose output
go test -v

# Run specific test
go test -run TestFunctionName

# Test with coverage
go test -cover
```

**When to use**: Ensuring code quality, continuous integration.

**Example**:

```go
// math_test.go
package main

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d; want 5", result)
    }
}
```

### `go test -bench`

Run benchmark tests.

```bash
go test -bench=.
go test -bench=BenchmarkFunction
```

**When to use**: Performance testing and optimization.

## Code Quality

### `go fmt`

Format Go source code.

```bash
# Format current directory
go fmt

# Format specific files
go fmt main.go

# Format all packages
go fmt ./...
```

**When to use**: Before committing code (or set up auto-format in your editor).

### `go vet`

Examine Go source code and report suspicious constructs.

```bash
go vet
go vet ./...
```

**When to use**: Code review, finding potential bugs.

### `go doc`

Show documentation for packages and symbols.

```bash
# Package documentation
go doc fmt
go doc net/http

# Function documentation
go doc fmt.Println
go doc json.Marshal
```

**When to use**: Learning about standard library or third-party packages.

## Package Management

### `go list`

List packages and modules.

```bash
# List current module
go list

# List all dependencies
go list -m all

# List available versions
go list -m -versions github.com/gorilla/mux
```

**When to use**: Understanding project structure, dependency management.

### `go mod why`

Explain why packages are needed.

```bash
go mod why github.com/gorilla/mux
```

**When to use**: Understanding dependency chains, cleaning up unused deps.

### `go mod graph`

Print module requirement graph.

```bash
go mod graph
```

**When to use**: Visualizing complex dependency relationships.

## Debugging and Profiling

### `go build -race`

Build with race condition detection.

```bash
go build -race
go test -race
```

**When to use**: Concurrent programming, finding race conditions.

### `go tool`

Access Go toolchain tools.

```bash
# CPU profiling
go tool pprof cpu.prof

# Memory profiling
go tool pprof mem.prof

# Assembly output
go tool objdump myapp
```

**When to use**: Performance analysis, debugging complex issues.

## Environment and Configuration

### `go env`

Print Go environment information.

```bash
# Show all environment variables
go env

# Show specific variable
go env GOPATH
go env GOOS GOARCH
```

**When to use**: Debugging environment issues, understanding configuration.

### `go env -w`

Set Go environment variables.

```bash
go env -w GOPROXY=direct
go env -w GO111MODULE=on
```

**When to use**: Configuring Go behavior, corporate environments.

### `go clean`

Remove object files and cached files.

```bash
# Clean current package
go clean

# Clean cache
go clean -cache

# Clean module cache
go clean -modcache
```

**When to use**: Freeing disk space, resolving build issues.

## Common Workflows

### Starting a New Project

```bash
mkdir myproject && cd myproject
go mod init myproject
echo 'package main; import "fmt"; func main() { fmt.Println("Hello!") }' > main.go
go run main.go
```

### Adding Dependencies

```bash
go get github.com/gorilla/mux
# Add import to your .go files
go mod tidy
```

### Building for Production

```bash
go test ./...
go vet ./...
go build -o myapp
```

### Cross-Platform Building

```bash
# Windows
GOOS=windows GOARCH=amd64 go build -o myapp.exe

# macOS
GOOS=darwin GOARCH=amd64 go build -o myapp-mac

# Linux ARM
GOOS=linux GOARCH=arm64 go build -o myapp-linux-arm64
```

## Tips for Beginners

1. **Always use modules**: Start every project with `go mod init`
2. **Format regularly**: Use `go fmt` or set up auto-formatting
3. **Test early**: Write tests alongside your code
4. **Read documentation**: Use `go doc` to understand packages
5. **Keep dependencies clean**: Run `go mod tidy` regularly

## Common Flags

- `-v`: Verbose output
- `-a`: Force rebuilding of packages
- `-o filename`: Output file name
- `-tags`: Build tags
- `-ldflags`: Linker flags
- `-gcflags`: Compiler flags

## Getting Help

- `go help [command]`: Built-in help
- [Official documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go.html)

## Next Steps

After mastering these commands:

1. Learn about Go's standard library
2. Explore popular frameworks (Gin, Echo, etc.)
3. Practice with real projects
4. Join the Go community

Remember: The best way to learn is by doing. Start with simple projects and gradually work your way up to more complex applications!
