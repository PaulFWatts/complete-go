package library

import (
	"os"
	"os/exec"
	"runtime"
)

// ClearScreen clears the console screen in a cross-platform manner.
// On Windows, it uses the 'cls' command.
// On Unix-like systems (Linux, macOS), it uses the 'clear' command.
func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
