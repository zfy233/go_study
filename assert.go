package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	rw := w.(io.ReadWriter) // success: *os.File has both Read and Write
	fmt.Println(rw)
	//w = new(ByteCounter)
	//rw = w.(io.ReadWriter) // panic: *ByteCounter has no Read method
}
