package modules

import (
	"fmt"
	"time"

	"github.com/vivafoxdirector/gomon/common"
)

type Cpu struct {
	CPU string `json:"cpu"`
}

func (c Cpu) Extract() {
	record := common.Record{Code: "cpu", Time: time.Now(), Value: map[string]string{"CPU": "CPU"}}
	fmt.Println(record)
}

func NewCPU() Module {
	cpu := Cpu{
		CPU: "aaaaa",
	}
	return Module(cpu)
}
