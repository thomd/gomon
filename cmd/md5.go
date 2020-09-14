package cmd

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func fileMd5(filePath string) (string, error) {
	var md5String string

	file, err := os.Open(filePath)
	if err != nil {
		return md5String, err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return md5String, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	md5String = hex.EncodeToString(hashInBytes)

	return md5String, nil
}
