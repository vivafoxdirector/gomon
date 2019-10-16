package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

type SystemInfo struct {
	Time            int64  `json:"timeslice"`
	Name            string `json:"hostname"`
	*SystemResource `json:"sr"`
}

type SystemResource struct {
	Cpu   float64 `json:"cpu"`
	Mem   float64 `json:"mem"`
	*Disk `json:"disk"`
}

type Disk struct {
	Total string `json:"total"`
	Free  string `json:"free"`
	Used  string `json:"used"`
}

// 정보 가져오는 주기는 3 sec로 한다
func printResourceInfo() {
	// Cpu채널 생성
	_c := make(chan []float64)
	for true {
		// 1. make Json Data
		fmt.Println(makeJsonValue(_c))

		// 2. Send to Server
	}
}

func makeJsonValue(_c chan []float64) string {
	// Cpu사용량 가져오기
	go cpuPercentage(_c)
	cpuP, _ := <-_c // 배열 채널은 이와같이 값을 가져오도록 한다.

	// 메모리 사용량 가져오기
	memP, _ := mem.VirtualMemory()
	//		fmt.Printf("cpu: %f%%, mem: %f%%\n", cpuP[0], memP.UsedPercent)

	// Disk 사용량 가져오기
	d, err := disk.Usage("/")
	check(err)

	v := &SystemInfo{
		time.Now().Unix(),
		getHostName(),
		&SystemResource{
			cpuP[0],
			memP.UsedPercent,
			&Disk{ // GIB 단위로 한다
				strconv.FormatUint(d.Total/1024/1024/1024, 10),
				strconv.FormatUint(d.Free/1024/1024/1024, 10),
				strconv.FormatUint(d.Used/1024/1024/1024, 10),
			},
		},
	}

	bytes, err := json.Marshal(v)
	check(err)
	return string(bytes)
}

func getHostName() string {
	name, _ := os.Hostname()
	return name
}

func printUsage(u *disk.UsageStat) {
	fmt.Println(u.Path + "\t" + strconv.FormatFloat(u.UsedPercent, 'f', 2, 64) + "% full.")
	fmt.Println("Total: " + strconv.FormatUint(u.Total/1024/1024/1024, 10) + " GiB")
	fmt.Println("Free:  " + strconv.FormatUint(u.Free/1024/1024/1024, 10) + " GiB")
	fmt.Println("Used:  " + strconv.FormatUint(u.Used/1024/1024/1024, 10) + " GiB")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func customSleep(s int) {
	time.Sleep(time.Duration(s) * time.Second)
}

// Cpu정보 가져오기(Interval사용으로 Go Routine사용)
func cpuPercentage(cpuP chan []float64) {
	duration := time.Duration(3) * time.Second
	cp, _ := cpu.Percent(duration, false)
	cpuP <- cp
}

func printCpuPersentage() {
	cpuP := make(chan float64)
	duration := time.Duration(3) * time.Second
	go func() {
		cp, _ := cpu.Percent(duration, false)
		cpuP <- cp[0]
	}()
	fmt.Printf("%f%%\n", <-cpuP)
}

func printCpu() {
	// CPU
	duration := time.Duration(3) * time.Second
	percentage, err := cpu.Percent(duration, true)
	if err != nil {
		return
	}
	fmt.Println("CPU%: ", percentage)
	cpuSum := float64(0)
	for _, per := range percentage {
		cpuSum = cpuSum + per
	}
	fmt.Println("CPUSUM: ", cpuSum)
	cpuAvg := uint32(cpuSum / float64(len(percentage)))
	fmt.Println("CPUAVG: ", cpuAvg)
}

func printCpuInfo() {
	// Cpu 정보 표시
	c, _ := cpu.Times(false)
	for _, v := range c {
		fmt.Println(v)
	}
	// Cpu 사용량 표시 (interval : 3second)
	duration := time.Duration(3) * time.Second
	for true {
		cp, _ := cpu.Percent(duration, false)
		fmt.Printf("%f%%\n", cp[0])
	}
}

func printMemInfo() {
	for true {
		v, _ := mem.VirtualMemory()
		// almost every return value is a struct
		fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
		time.Sleep(time.Duration(3) * time.Second)
		// convert to JSON. String() is also implemented
		//fmt.Println(v)
	}
}

func main() {
	printResourceInfo()
	//printCpu()
}
