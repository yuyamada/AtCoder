package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	d := io.NextInt()
	for i := 0; i < d; i++ {
		io.Println(1)
	}
}
