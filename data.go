package main

import (
	"log"
	"os"
)

func ensureDir() {
	dirErr := os.Mkdir("./data", os.ModePerm)

	if dirErr != nil && !os.IsExist(dirErr) {
		log.Fatal(dirErr)
	}
}

func ensureFile(filePath string) {
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
