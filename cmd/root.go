package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/spf13/cobra"
)

var (
	VERSION        = "0.2.0"
	monitoringPath string
	ignoreDirs     []string
	pid            int
	program        string
)

func gomon(cmd *cobra.Command, args []string) {
	program = args[0]
	monitoringPath = "./"

	filepath.Walk(monitoringPath, filesToWatch(storeHash))

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
	Use: "gomon",
	Long: `
Gomon will monitor for any changes in your Go source code and automatically restart your app`,
	Args:    cobra.MinimumNArgs(1),
	Version: VERSION,
	Run:     gomon,
}

func Execute() {
	rootCmd.Flags().StringSliceVarP(&ignoreDirs, "ignore", "i", []string{}, "ignore directories")
	rootCmd.SetVersionTemplate(`gomon version: {{.Version}}`)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
