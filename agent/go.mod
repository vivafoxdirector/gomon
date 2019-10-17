module bitbucket.org/vivafoxdirector/gomon/agent

go 1.12

replace bitbucket.org/vivafoxdirector/gomon/common => ../common

replace bitbucket.org/vivafoxdirector/gomon/agent/modules => ./modules

require (
	bitbucket.org/vivafoxdirector/gomon/common v0.0.0-00010101000000-000000000000
	github.com/shirou/gopsutil v2.18.12+incompatible
	golang.org/x/sys v0.0.0-20190804053845-51ab0e2deafa // indirect
)
