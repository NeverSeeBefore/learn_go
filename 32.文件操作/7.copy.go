package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	var err error
	err = ioutilCp("./note.md", "./note.md.ioutil.cp")
	if err != nil {
		fmt.Println("ioutilCp failed: ", err)
	}

	err = fileCp("./note.md", "./note.md.file.cp")
	if err != nil {
		fmt.Println("fileCp failed: ", err)
	}

}

func ioutilCp(origin, target string) error {
	byteStr, err := ioutil.ReadFile(origin)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(target, byteStr, 0666)
	if err != nil {
		return err
	}
	return nil
}

func fileCp(origin, target string) error {
	var err error
	originFile, err := os.Open(origin)
	defer originFile.Close()
	if err != nil {
		return err
	}
	targetFile, err := os.OpenFile(target, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	defer targetFile.Close()
	if err != nil {
		return err
	}

	var tempSlice = make([]byte, 128)
	for {
		n, err := originFile.Read(tempSlice)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		targetFile.Write(tempSlice[:n])
	}

	return nil
}
