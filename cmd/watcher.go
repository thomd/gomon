package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func fileWatcher() {
	for {
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

			fileHash, err := fileMd5(path)
			if err != nil {
				panic(`could not calculate hash for ` + path)
			}

			if _, ok := fileHashes[path]; !ok {
				fileHashes[path], _ = fileMd5(path)

			} else if fileHashes[path] != fileHash {
				fileHashes[path] = fileHash

				fmt.Println(`file changed`, path, ` . Restarting web server`)
				//cmd.Process.Kill()
				//cmd.Run()
			}
			return nil
		})

		time.Sleep(100 * time.Millisecond)
	}
}
