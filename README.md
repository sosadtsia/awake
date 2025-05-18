# awake

A command-line tool for macOS that prevents your Mac from sleeping.

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

Show debug information:
```
awake -d
```

## Implementation

This tool uses the built-in macOS `caffeinate` command to prevent your Mac from sleeping. The tool provides:

- Time-limited operation (automatically exit after a specified duration)
- Quiet mode for background operation
- Debug logging for troubleshooting
- Proper signal handling for clean shutdown

## How It Works

1. Parses command-line options to determine behavior
2. Runs caffeinate with appropriate flags:
   - `-d` to prevent display sleep
   - `-i` to prevent system idle sleep
   - `-t` when a time limit is specified
3. Properly terminates the caffeinate process when:
   - The specified duration elapses
   - The user presses Ctrl+C
   - The process receives a termination signal

## Development

### Pre-commit Hooks

This project uses git pre-commit hooks to ensure code quality. The hooks:

1. Check if golangci-lint is installed
2. Run linting with golangci-lint
3. Run Go tests

When you clone the repository, enable the pre-commit hook:

```bash
# Make the pre-commit hook executable
chmod +x .git/hooks/pre-commit
```

If you don't have golangci-lint installed, you can install it with:

```bash
# macOS
brew install golangci-lint

# Other platforms
# See https://golangci-lint.run/usage/install/
```
