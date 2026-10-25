package main

import (
	"os"
	"path/filepath"
)

// panic is when sometihng goes wrong --> used to:
// 		fail in errors that should not occur in normal execution
// 		fail in case of a error to which we are not prepared to handle

func main() {
	panic("a problem")

	// one usage pattern is to fail when we didnt expect an error

	path := filepath.Join(os.TempDir(), "file")
	_, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	// AFTER A PANIC THE PROGRAM WILL CRASH LIKE IN RUST (the program exists)

	// in go, in functions is common to indicate a error return type (similar to Rust Result enum)
}
