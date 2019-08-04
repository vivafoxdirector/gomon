package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
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

func main() {
	agentStart()
}

// 정보 가져오는 주기는 3 sec로 한다
func agentStart() {
	// Cpu채널 생성
	_c := make(chan []float64)
	for true {
		// 1. make Json Data
		fmt.Println(makeJsonValue(_c))
		// 2. Send to Server

	}
}

func sendToServer(json string) {
	conn, _ := net.Dial("tcp", "localhost:9999")
	request, _ := http.NewRequest("GET", "http://localhost:9999", nil)
	request.Write(conn)
	response, _ := http.ReadResponse(bufio.NewReader(conn), request)
	dump, _ := httputil.DumpResponse(response, true)
	fmt.Println(string(dump))
}

func makeJsonValue(_c chan []float64) string {
	// Cpu사용량 가져오기
	go cpuPercentage(_c)
	cpuP, _ := <-_c // 배열 채널은 이와같이 값을 가져오도록 한다.

	// 메모리 사용량 가져오기
	memP, _ := mem.VirtualMemory()
	//		fmt.Printf("cpu: %f%%, mem: %f%%\n", cpuP[0], memP.UsedPercent)

	// Disk 사용량 가져오기
	d, _ := disk.Usage("/")
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Cpu정보 가져오기(Interval사용으로 Go Routine사용)
func cpuPercentage(cpuP chan []float64) {
	duration := time.Duration(3) * time.Second
	cp, _ := cpu.Percent(duration, false)
	cpuP <- cp
}
