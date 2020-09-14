package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "gomon",
	Long: `Gomon will monitor for any changes in your source code and automatically restart your app`,
	Run: func(cmd *cobra.Command, args []string) {
		monitoringPath := "./"
		skipDirectory := ".git"
		fileHashes := make(map[string]string)
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
		for key, value := range fileHashes {
			fmt.Printf("%s = %s\n", key, value)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
