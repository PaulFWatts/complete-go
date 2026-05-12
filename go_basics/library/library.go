package library

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// ClearScreen clears the console screen in a cross-platform manner.
// On Windows, it uses the "cls" command.
// On Unix-like systems (Linux, macOS), it uses the "clear" command.
func ClearScreen() error {
	const ansiClear = "\033[H\033[2J\033[3J"

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("clear screen: %w", err)
		}
		return nil
	}

	if term := os.Getenv("TERM"); term != "" && term != "dumb" {
		cmd = exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err == nil {
			return nil
		}
	}

	if _, err := fmt.Fprint(os.Stdout, ansiClear); err == nil {
		return nil
	}

	// Last-resort fallback for consoles that do not honor ANSI sequences.
	if _, err := fmt.Fprint(os.Stdout, strings.Repeat("\n", 80)); err != nil {
		return fmt.Errorf("clear screen: %w", err)
	}
	return nil
}
