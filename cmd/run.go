package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

var firstCall bool

func init() {
	firstCall = true
}

func runProgram() int {
	cmd := exec.Command("go", "run", program)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Start()

	if firstCall {
		fmt.Println("Start")
		firstCall = false
	} else {
		fmt.Println("Restart")
	}

	id, _ := syscall.Getpgid(cmd.Process.Pid)
	return id
}
