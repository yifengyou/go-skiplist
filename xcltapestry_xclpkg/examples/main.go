package main

import (
	"fmt"
	"github.com/xcltapestry/xclpkg"
)


func main() {
	slt := algorithm.NewSkipList()
	for i := 100; i > 0; i-- {
		slt.Insert(i)
	}
	slt.PrintSkipList()
	slt.Search(15)
	slt.Search(93)
	slt.Remove(93)
	slt.PrintSkipList()
	fmt.Println("Done!")
}
