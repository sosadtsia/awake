// awake is a CLI tool that prevents your Mac from sleeping
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

// Version information
const (
	appName    = "awake"
	appVersion = ""
)

// CLI options
type options struct {
	quiet      bool
	debug      bool
	version    bool
	help       bool
	duration   time.Duration
	background bool
}

// runInBackground starts a detached process and exits the current one
func runInBackground(opts options, logger *logger) {
	// Only do this if we're not already a background process
	if os.Getenv("AWAKE_BACKGROUND") != "1" {
		logger.Debug("Starting background mode")

		// Get the current executable path
		executable, err := os.Executable()
		if err != nil {
			logger.Error("Error getting executable path: %v", err)
			os.Exit(1)
		}

		// Prepare command for background execution
		args := []string{}
		// Copy all args except the background flag
		for _, arg := range os.Args[1:] {
			if arg != "-background" && arg != "-b" {
				args = append(args, arg)
			}
		}

		// Add the quiet flag if not already present
		hasQuiet := false
		for _, arg := range args {
			if arg == "-q" || arg == "-quiet" {
				hasQuiet = true
				break
			}
		}
		if !hasQuiet {
			args = append(args, "-q")
		}

		logger.Debug("Spawning background process with args: %v", args)

		// Create a new command with the same binary and args
		cmd := exec.Command(executable, args...)
		// Set environment variable to prevent infinite recursion
		cmd.Env = append(os.Environ(), "AWAKE_BACKGROUND=1")
		// Detach process
		cmd.Stdin = nil
		cmd.Stdout = nil
		cmd.Stderr = nil

		if err := cmd.Start(); err != nil {
			logger.Error("Failed to start background process: %v", err)
			os.Exit(1)
		}

		logger.Info("Awake is now running in the background (PID: %d)", cmd.Process.Pid)
		if opts.duration > 0 {
			logger.Info("Mac will stay awake for %s", opts.duration)
			logger.Info("Use 'pkill awake' to stop it early if needed")
		} else {
			logger.Info("Mac will stay awake until you run 'pkill awake'")
		}
		os.Exit(0)
	}
}

func main() {
	// Parse command line options
	opts := parseFlags()

	// Handle special flags first
	if opts.version {
		fmt.Printf("%s v%s\n", appName, appVersion)
		return
	}

	if opts.help {
		flag.Usage()
		return
	}

	// Initialize logger based on options
	logger := newLogger(opts.quiet, opts.debug)

	// Handle background mode if requested
	if opts.background {
		runInBackground(opts, logger)
		// if we're still here, we're the child process
	}

	// Check for caffeinate
	path, err := exec.LookPath("caffeinate")
	if err != nil {
		logger.Error("caffeinate command not found. This tool requires macOS.")
		os.Exit(1)
	}

	logger.Debug("Found caffeinate at: %s", path)
	logger.Info("Starting awake - preventing sleep on your Mac")

	// Setup caffeinate arguments
	args := []string{"-d", "-i"} // Prevent display sleep and system idle sleep

	// Add time limit if specified
	if opts.duration > 0 {
		args = append(args, "-t", fmt.Sprintf("%d", int(opts.duration.Seconds())))
		logger.Info("Mac will stay awake for %s", opts.duration)
	} else {
		logger.Info("Mac will stay awake until you press Ctrl+C")
	}

	// Start caffeinate
	cmd := exec.Command("caffeinate", args...)
	if err := cmd.Start(); err != nil {
		logger.Error("Failed to start caffeinate: %v", err)
		os.Exit(1)
	}

	// If running with duration, wait for completion
	if opts.duration > 0 {
		logger.Debug("Waiting for %s to complete", opts.duration)
		go func() {
			if err := cmd.Wait(); err != nil {
				logger.Debug("caffeinate process ended: %v", err)
			}
			logger.Info("Duration completed, exiting")
			os.Exit(0)
		}()
	}

	// Set up channel for signal handling
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Wait for signal
	<-c

	// Kill caffeinate process
	if err := cmd.Process.Kill(); err != nil {
		logger.Error("Failed to stop caffeinate: %v", err)
	} else {
		logger.Info("Stopped caffeinate - Mac can sleep normally now")
	}
}

// Parse command line flags
func parseFlags() options {
	var opts options

	flag.BoolVar(&opts.quiet, "quiet", false, "Suppress all output")
	flag.BoolVar(&opts.quiet, "q", false, "Suppress all output (shorthand)")

	flag.BoolVar(&opts.debug, "debug", false, "Enable debug logging")
	flag.BoolVar(&opts.debug, "d", false, "Enable debug logging (shorthand)")

	flag.BoolVar(&opts.version, "version", false, "Show version information")
	flag.BoolVar(&opts.version, "v", false, "Show version information (shorthand)")

	flag.BoolVar(&opts.help, "help", false, "Show help information")
	flag.BoolVar(&opts.help, "h", false, "Show help information (shorthand)")

	flag.DurationVar(&opts.duration, "time", 0, "Duration to prevent sleep (e.g. 1h30m)")
	flag.DurationVar(&opts.duration, "t", 0, "Duration to prevent sleep (shorthand)")

	flag.BoolVar(&opts.background, "background", false, "Run in background mode")
	flag.BoolVar(&opts.background, "b", false, "Run in background mode (shorthand)")

	// Create custom usage message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", appName)
		fmt.Fprintf(os.Stderr, "A tool that prevents your Mac from sleeping using the caffeinate command.\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  %s                         # Prevent sleep until Ctrl+C\n", appName)
		fmt.Fprintf(os.Stderr, "  %s -t 2h                   # Prevent sleep for 2 hours\n", appName)
		fmt.Fprintf(os.Stderr, "  %s -q -t 30m               # Quietly prevent sleep for 30 minutes\n", appName)
		fmt.Fprintf(os.Stderr, "  %s -b                      # Run in background indefinitely (use 'pkill awake' to stop)\n", appName)
		fmt.Fprintf(os.Stderr, "  %s -b -t 2h                # Run in background for 2 hours\n", appName)
		fmt.Fprintf(os.Stderr, "  %s -b -t 2h -q             # Run in background for 2 hours, quietly\n", appName)
	}

	flag.Parse()
	return opts
}

// Logger handles formatted output based on verbosity settings
type logger struct {
	quiet bool
	debug bool
}

// Create a new logger
func newLogger(quiet, debug bool) *logger {
	return &logger{
		quiet: quiet,
		debug: debug,
	}
}

// Info logs informational messages
func (l *logger) Info(format string, args ...interface{}) {
	if !l.quiet {
		fmt.Printf(format+"\n", args...)
	}
}

// Debug logs debug messages (only in debug mode)
func (l *logger) Debug(format string, args ...interface{}) {
	if l.debug && !l.quiet {
		fmt.Printf("[DEBUG] "+format+"\n", args...)
	}
}

// Error logs error messages (even in quiet mode)
func (l *logger) Error(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "ERROR: "+format+"\n", args...)
}
