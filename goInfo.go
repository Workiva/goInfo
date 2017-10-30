package goInfo

import (
	"fmt"
)

type GoInfoObject struct {
	GoOS         string `json:"baseOs"`
	Kernel       string `json:"kernel"`
	Core         string `json:"core"`
	Platform     string `json:"platform"`
	OS           string `json:"os"`
	Hostname     string `json:"hostname"`
	CPUs         int    `json:"cpus"`
	Distribution string `json:"distribution"`
	Name         string `json:"name"`
}

func (gi *GoInfoObject) VarDump() {
	fmt.Println("GoOS:", gi.GoOS)
	fmt.Println("Kernel:", gi.Kernel)
	fmt.Println("Core:", gi.Core)
	fmt.Println("Platform:", gi.Platform)
	fmt.Println("OS:", gi.OS)
	fmt.Println("Hostname:", gi.Hostname)
	fmt.Println("CPUs:", gi.CPUs)
	if gi.Kernel == "Linux" {
		fmt.Println("Distribution version:", gi.Distribution)
		fmt.Println("Distribution name:", gi.Name)
	}
}
func (gi *GoInfoObject) Dist() {
	if gi.Kernel == "Linux" {
		fmt.Println("Distribution version:", gi.Distribution)
	}
}
func (gi *GoInfoObject) DistName() {
	if gi.Kernel == "Linux" {
		fmt.Println("Distribution name:", gi.Name)
	}
}
func (gi *GoInfoObject) String() string {
	return fmt.Sprintf("GoOS:%v,Kernel:%v,Core:%v,Platform:%v,OS:%v,Hostname:%v,CPUs:%v, Distribution:%v", gi.GoOS, gi.Kernel, gi.Core, gi.Platform, gi.OS, gi.Hostname, gi.CPUs, gi.Distribution)
}
