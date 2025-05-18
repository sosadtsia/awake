# awake

A command-line tool for macOS that prevents your Mac from sleeping.

## Requirements

- macOS (relies on the built-in `caffeinate` command)
- Go 1.18 or higher (for building from source)

## Installation

### Option 1: Manual Installation

1. Clone the repository
   ```
   git clone https://github.com/sosadtsia/awake.git
   cd awake
   ```

2. Build with Go
   ```
   go build -o awake
   ```

3. Install to a directory in your PATH (choose one):
   ```
   # Using Homebrew location (recommended for M1/M2 Macs)
   cp awake /opt/homebrew/bin/

   # Using traditional location (may require sudo)
   sudo cp awake /usr/local/bin/

   # Using user bin directory
   mkdir -p ~/.local/bin
   cp awake ~/.local/bin/
   # Ensure ~/.local/bin is in your PATH by adding this to your shell profile:
   # export PATH="$HOME/.local/bin:$PATH"
   ```

### Option 2: Using Go Install
```
go install github.com/sosadtsia/awake@latest
```
This automatically installs the binary to your `$GOPATH/bin` directory (usually `~/go/bin`). Ensure this directory is in your PATH.

### Option 3: Download Pre-built Binary

You can download pre-built binaries for macOS from the [Releases](https://github.com/sosadtsia/awake/releases) page. Choose from:

- **darwin-amd64**: For Intel Macs
- **darwin-arm64**: For Apple Silicon Macs (M1/M2/M3)
- **darwin-universal**: Universal binary that works on both Intel and Apple Silicon Macs

Compressed versions (`.gz`) of each binary are also available for faster downloads.

After downloading, make the binary executable and move it to a location in your PATH:

```bash
# For regular binaries
chmod +x awake-v*-darwin-*
mv awake-v*-darwin-* /usr/local/bin/awake

# For compressed (.gz) binaries
gunzip awake-v*-darwin-*.gz
chmod +x awake-v*-darwin-*
mv awake-v*-darwin-* /usr/local/bin/awake
```

### Verify Installation

To verify installation:
```
which awake
awake -v
```

## Usage

```
awake [options]
```

### Options

| Option | Shorthand | Description |
|--------|-----------|-------------|
| `-quiet` | `-q` | Suppress all output |
| `-debug` | `-d` | Enable debug logging |
| `-version` | `-v` | Show version information |
| `-help` | `-h` | Show help information |
| `-time DURATION` | `-t DURATION` | Set a duration to prevent sleep (e.g., "2h", "30m", "1h30m") |
| `-background` | `-b` | Run in background mode (detached from terminal) |

### Examples

Prevent sleep until manually stopped with Ctrl+C:
```
awake
```

Prevent sleep for 2 hours:
```
awake -t 2h
```

Quietly prevent sleep for 30 minutes:
```
awake -q -t 30m
```

Run in background mode (detached from terminal):
```
awake -b
```

Run in background mode for 2 hours:
```
awake -b -t 2h
```

Show debug information:
```
awake -d
```

## Implementation

This tool uses the built-in macOS `caffeinate` command to prevent your Mac from sleeping. The tool provides:

- Time-limited operation (automatically exit after a specified duration)
- Background mode for detached operation
- Quiet mode for no terminal output
- Debug logging for troubleshooting
- Proper signal handling for clean shutdown

## How It Works

1. Parses command-line options to determine behavior
2. When run in background mode, it spawns a detached copy of itself and exits
3. Runs caffeinate with appropriate flags:
   - `-d` to prevent display sleep
   - `-i` to prevent system idle sleep
   - `-t` when a time limit is specified
4. Properly terminates the caffeinate process when:
   - The specified duration elapses
   - The user presses Ctrl+C (in foreground mode)
   - The process receives a termination signal

## Background Mode

When run with the `-b` or `-background` flag, awake will:

1. Create a detached copy of itself with the same parameters (minus the `-b` flag)
2. Print the PID of the background process
3. Exit the parent process, leaving the detached process running
4. To stop a background process, use `pkill awake`

## Versioning

This project follows [Semantic Versioning](https://semver.org/) (SemVer).

- **Major version** increments denote incompatible API changes
- **Minor version** increments denote added functionality in a backward-compatible manner
- **Patch version** increments denote backward-compatible bug fixes

Release tags follow the `vX.Y.Z` format (e.g., `v1.0.0`).

## Releases

New releases are automatically created via GitHub Actions when commits following the [Conventional Commits](https://www.conventionalcommits.org/) format are pushed to the main branch. For example:

- `feat: ...` - Triggers a minor version bump
- `fix: ...` - Triggers a patch version bump
- `feat!: ...` or including `BREAKING CHANGE:` in commit message - Triggers a major version bump

See [CONTRIBUTING.md](CONTRIBUTING.md) for more details on the commit message format.

## Development

For detailed development instructions, please see [DEVELOPMENT.md](DEVELOPMENT.md).

### Task Runner

This project uses [Task](https://taskfile.dev) for managing development tasks:

```bash
# Install Task
go install github.com/go-task/task/v3/cmd/task@latest

# Run tasks
task build    # Build the application
task test     # Run tests
task lint     # Run linter
task hooks    # Setup git hooks
task clean    # Clean build artifacts
```

### Pre-commit Hooks

This project uses git pre-commit hooks to ensure code quality. The hooks:

1. Check if golangci-lint is installed
2. Run linting with golangci-lint
3. Run Go tests

#### Installing the hooks

Set up the git hooks using Task:

```bash
# Using Task
task hooks
```

This will configure git to use the hooks in the `.githooks` directory.

#### Requirements

The pre-commit hook requires golangci-lint. If you don't have it installed, you can install it with:

```bash
# macOS
brew install golangci-lint
```
