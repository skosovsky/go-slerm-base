package main

import (
	"bytes"
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	copyFile, err := os.OpenFile("copy.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = copyFile.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	_, err = copyFile.Write(file)
	if err != nil {
		log.Println(err)
	}

	log.Println("result:", checkFiles("test.txt", "copy.txt"))
}

func checkFiles(dst string, src string) bool {
	dstFile, err := os.ReadFile(dst)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(dstFile)
	srcFile, err := os.ReadFile(src)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(srcFile)

	return bytes.Equal(dstFile, srcFile)
}
