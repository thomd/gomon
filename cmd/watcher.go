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

func hashFiles(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() && info.Name() == skipDirectory {
		return filepath.SkipDir
	}
	if info.IsDir() {
		return nil
	}
	fileHashes[path], err = fileMd5(path)
	if err != nil {
		panic(`could not calculate hash for ` + path)
	}

	return nil
}

func restartOnFileChange(path string, info os.FileInfo, err error) error {
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

	if hash, ok := fileHashes[path]; !ok || hash != fileHash {
		fileHashes[path] = fileHash
		fmt.Printf(yellow("[gomon] file '%s' changed\n"), path)
		syscall.Kill(-pid, 15)
		pid = runProgram()
	}

	return nil
}

func fileWatcher() {
	for {
		filepath.Walk(monitoringPath, restartOnFileChange)
		time.Sleep(100 * time.Millisecond)
	}
}
