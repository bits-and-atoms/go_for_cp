package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	fmt.Fprintln(out, n)
}