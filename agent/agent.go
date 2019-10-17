package main

import (
	"flag"
	"fmt"

	"bitbucket.org/vivafoxdirector/gomon/agent/modules"
)

/**
 * ## AGENT ##
 * 1. 입력
 *
 * 2. 출력
 *    - http
 *    - screen output
 */
func main() {
	var (
		i = flag.String("i", "localhost", "server ip")
		p = flag.Int("p", 8080, "server port")
		o = flag.String("o", "http", "output format")
	)
	flag.Parse()
	fmt.Println(*i, *p, *o)

	cpu := modules.NewCPU()

	fmt.Println(cpu)

}
