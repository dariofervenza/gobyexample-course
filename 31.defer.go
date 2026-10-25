package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// defer is used to ensure a function is called later in the program, often for cleanups
// its like the finally in python

func main() {

	// common examples are closing a file after writing or closing a db connection

	path := filepath.Join(os.TempDir(), "defer.txt")
	f := createFile(path)
	defer closeFile(f)
	writeFile(f)

}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	// use Fprintln to write into an io buffer instead of the stdout
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	// deferred calls also need to check errors, its important
	if err != nil {
		panic(err)
	}
}
