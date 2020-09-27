package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

var firstCall = true

func runProgram() int {
	cmd := exec.Command("go", "run", program)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Start()

	pid, _ := syscall.Getpgid(cmd.Process.Pid)

	if firstCall {
		fmt.Printf(yellow("[gomon] %s\n"), VERSION)
		fmt.Printf(yellow("[gomon] monitoring files: *%s\n"), fileExtension)
		if len(ignoreDirs) > 0 {
			fmt.Printf(yellow("[gomon] ignoring folders: %s\n"), strings.Join(ignoreDirs[:], ", "))
		}
		fmt.Printf(green("[gomon] starting 'go run %s' (pid:%d)\n"), program, pid)
		firstCall = false
	} else {
		fmt.Printf(green("[gomon] restarting 'go run %s' (pid:%d)\n"), program, pid)
	}

	return pid
}
