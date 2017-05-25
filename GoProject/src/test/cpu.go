package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpu_num := runtime.NumCPU()
	fmt.Println(cpu_num)
}
