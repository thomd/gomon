package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/spf13/cobra"
)

var monitoringPath string
var skipDirectory string
var pid int
var program string

func gomon(cmd *cobra.Command, args []string) {
	program = args[0]
	monitoringPath = "./"
	skipDirectory = ".git"

	filepath.Walk(monitoringPath, hashFiles)

	done := make(chan bool)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	pid = runProgram()

	go func() {
		<-signalChan
		syscall.Kill(-pid, 15)
		done <- true
	}()

	go fileWatcher()

	<-done
}

var rootCmd = &cobra.Command{
	Use:  "gomon",
	Long: `gomon will monitor for any changes in your source code and automatically restart your app`,
	Args: cobra.MinimumNArgs(1),
	Run:  gomon,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
