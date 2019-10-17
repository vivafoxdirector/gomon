package common

import "fmt"
import "strings"

type gomonProperties map[string]string

func gomonPropertiesUtil() {
	var config = properties{}

	if index := strings.Index(prop, "="); index >= 0 {
		if key := strings.TrimSpace(prop[:index]); len(key) > 0 {
			if len(prop) > index+1 {
				val := strings.TrimSpace(prop[index+1:])
				config[key] = val
			}
		}
	}
}
