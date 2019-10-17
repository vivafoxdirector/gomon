package modules

import (
	"fmt"
	"time"

	"bitbucket.org/vivafoxdirector/gomon/common"
)

type Cpu struct {
	CPU string `json:"cpu"`
}

func (c Cpu) Extract() {
	record := common.Record{Code: "cpu", Time: time.Now(), "CPU": "CPU"}
	fmt.Println(record)
}

func NewCPU() Module {
	cpu := Cpu{
		CPU: "",
	}
	return Module(cpu)
}
