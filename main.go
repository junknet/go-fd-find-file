package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	file := os.Args[1]
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.Contains(path, file) {
				out := strings.ReplaceAll(path, file, "\033[1;31;40m"+file+"\033[0m")
				fmt.Println(out)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
