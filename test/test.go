package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	printCpuInfo()
	//printMemInfo()
}

func printCpuInfo() {
	c, _ := cpu.Times(false)
	duration := time.Duration(10) * time.Microsecond
	cf, _ := cpu.Percent(duration, true)
	for _, v := range c {
		fmt.Println(v)
	}

	fmt.Println(cf)
}
func printMemInfo() {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)
}
