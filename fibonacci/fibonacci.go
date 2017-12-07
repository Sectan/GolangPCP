package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func fibonacci(n int) *big.Int {
	fn := make(map[int]*big.Int)

	if n == 0 {
		var f = big.NewInt(0)
		fn[n] = f
		return fn[n]
	}

	for i := 0; i <= n; i++ {
		var f = big.NewInt(0)
		if i <= 2 {
			f.SetUint64(1)
		} else {
			f = f.Add(fn[i-1], fn[i-2])
		}
		fn[i] = f
	}
	return fn[n]
}

func main() {
	for _, s := range os.Args[1:] {
		n, e := strconv.Atoi(s)
		if e != nil {
			panic(e)
		}
		fmt.Printf("%s\n", fibonacci(n))
	}
}
