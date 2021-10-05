package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
)

func CheckMd5(md5Hash string) (err error) {
	file, err := os.Open("file.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() == md5Hash {
			return fmt.Errorf("already exist")
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func SaveMd5Hash(md5Hash string) error {
	file, err := os.OpenFile("file.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.WriteString(md5Hash + "\n"); err != nil {
		return err
	}
	return nil
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
