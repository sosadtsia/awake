# Development Guide

This document provides instructions for developers contributing to the awake project.

## Setup

1. Clone the repository
   ```
   git clone https://github.com/yourusername/awake.git
   cd awake
   ```

2. Install Task (if not already installed)
   ```
   go install github.com/go-task/task/v3/cmd/task@latest
   ```

3. Install golangci-lint (if not already installed)
   ```
   # macOS
   brew install golangci-lint

   # Other platforms
   # See https://golangci-lint.run/usage/install/
   ```

4. Set up git hooks
   ```
   task hooks
   ```

## Development Workflow

### Building

```
task build
```

This will create the `awake` binary in the current directory.

### Running Tests

```
task test
```

### Running the Linter

```
task lint
```

### Running Pre-commit Checks Manually

You can manually run the same checks that the pre-commit hook runs:

```
task lint   # Run the linter
task test   # Run the tests
```

Alternatively, you can run both with a single command:

```
task precommit
```

## Project Features

### Background Mode

The application supports running in the background (detached from the terminal) using the `-b` or `-background` flag. When this flag is used:

1. The application spawns a new detached process with the same parameters (minus the background flag)
2. The parent process exits, leaving the child running in the background
3. The background process can be stopped using `pkill awake`

Implementation details:
- The background mode uses the `os/exec` package to spawn a detached process
- It sets the `AWAKE_BACKGROUND=1` environment variable to prevent recursive spawning
- The detached process has stdin, stdout, and stderr set to nil to fully detach from the terminal

## Git Hooks

### Pre-commit Hook

The pre-commit hook runs automatically before each commit and:

1. Checks if Task is installed
2. Checks if golangci-lint is installed
3. Runs `task lint` to check code quality
4. Runs `task test` to verify tests pass

If any of these checks fail, the commit will be aborted.

## Project Structure

- `main.go` - Main application code
- `main_test.go` - Tests for the application
- `.githooks/` - Git hooks
- `Taskfile.yml` - Task definitions for development workflows
