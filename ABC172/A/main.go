package main

import (
	"os"

	"github.com/youyouryu/AtCoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	a := io.NextInt()
	io.Println(a + a*a + a*a*a)
}