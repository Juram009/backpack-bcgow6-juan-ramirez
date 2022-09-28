package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("Hola")
	_, err := io.Copy(os.Stdout, reader)
	if err != nil {
		panic(err)
	}
	//b, err := io,ReadAll(reader)
	io.WriteString(os.Stdout, "HOLA :D")
}
