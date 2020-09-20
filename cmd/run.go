package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/gookit/color"
)

var firstCall = true
var green = color.FgGreen.Render
var yellow = color.FgYellow.Render

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
		fmt.Printf(yellow("[gomon] monitoring '%s*.*'\n"), monitoringPath)
		fmt.Printf(yellow("[gomon] excluding '%s%s'\n"), monitoringPath, skipDirectory)
		fmt.Printf(green("[gomon] starting 'go run %s' (pid:%d)\n"), program, pid)
		firstCall = false
	} else {
		fmt.Printf(green("[gomon] restarting 'go run %s' (pid:%d)\n"), program, pid)
	}

	return pid
}
