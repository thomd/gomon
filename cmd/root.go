package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var fileHashes map[string]string
var monitoringPath string
var skipDirectory string
var cmd *exec.Cmd

var rootCmd = &cobra.Command{
	Use:  "gomon",
	Long: `gomon will monitor for any changes in your source code and automatically restart your app`,
	Run: func(cmd *cobra.Command, args []string) {
		monitoringPath = "./"
		skipDirectory = ".git"

		fileHashes = make(map[string]string)
		filepath.Walk(monitoringPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() && info.Name() == skipDirectory {
				return filepath.SkipDir
			}
			if info.IsDir() {
				return nil
			}
			fileHashes[path], _ = fileMd5(path)
			return nil
		})

		go fileWatcher()

		// Run the web server
		fmt.Println(`Started server`)
		//cmd = exec.Command(`go`, `run`, `web/main.go`)
		//cmd.Run()

		// Create a channel and wait on it. This is here so the main thread exit
		doneChannel := make(chan bool)
		_ = <-doneChannel
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
