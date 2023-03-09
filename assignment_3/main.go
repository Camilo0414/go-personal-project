package main

import (
	"fmt"
	"io"
	"os"
)

type fileReader struct{}

func (fileReader) Read(b []byte) (int, error) {
	return len(b), nil
}

func main() {
	fn := os.Args[1]
	file, err := os.Open(fn)
	if err != nil {
		fmt.Println("This is fucked up:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	contents, err := os.ReadFile(os.Args[1])
// 	if err != nil {
// 		fmt.Println("File reading error", err)
// 		return
// 	}
// 	fmt.Println(string(contents))
// }
