package main

import (
	"fmt"
	"os"
)

func main() {
	var err error
	err = os.Mkdir("./aaa", 0777)
	fmt.Println(err)
	err = os.MkdirAll("./aaa/bbb/ccc", 0777)
	fmt.Println(err)
	err = os.Rename("./aaa", "./xxx")
	fmt.Println(err)
	err = os.Remove("./xxx")
	fmt.Println(err)
	err = os.RemoveAll("./xxx")
	fmt.Println(err)
	err = os.Rename("./note.md.ioutil copy.cp", "note.md.ioutil copy.cp.rename")
	fmt.Println(err)
}
