package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func checkte(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	f, err := ioutil.TempFile("", "sample")
	checkte(err)

	fmt.Println("Temp file name:", f.Name())

	defer os.Remove(f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	checkte(err)

	dname, err := ioutil.TempDir("", "sampledir")
	checkte(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
	checkte(err)
}
