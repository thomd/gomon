package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

var fileHashes map[string]string

func init() {
	fileHashes = make(map[string]string)
}

func filesToWatch(callback func(path string, fileHash string)) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			for _, dir := range ignoreDirs {
				if dir == info.Name() {
					return filepath.SkipDir
				}
			}
			return nil
		}
		if filepath.Ext(path) != fileExtension {
			return nil
		}
		fileHash, err := fileMd5(path)
		if err != nil {
			panic(`could not calculate hash for ` + path)
		}
		callback(path, fileHash)
		return nil
	}
}

func storeHash(path, fileHash string) {
	fileHashes[path] = fileHash
}

func restartOnFileChange(path, fileHash string) {
	if hash, ok := fileHashes[path]; !ok || hash != fileHash {
		action := "changed"
		if !ok {
			action = "added"
		}
		fmt.Printf(yellow("[gomon] file '%s' %s\n"), path, action)
		fileHashes[path] = fileHash
		syscall.Kill(-pid, 15)
		pid = runProgram()
	}
}

func fileWatcher() {
	for {
		filepath.Walk(monitoringPath, filesToWatch(restartOnFileChange))
		time.Sleep(100 * time.Millisecond)
	}
}
