package main

import (
	"fmt"
	"linkedList"
)

func main() {
	list := linkedList.NewDoubleList(1)
	list.Append(2)
	list.Append(4)
	list.Remove(4)
	fmt.Println(list.ToSlice())
}