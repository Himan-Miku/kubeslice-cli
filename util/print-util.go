package util

import (
	"fmt"
	"os"
)

const (
	Cross = string(rune(0x274c))
	Tick  = string(rune(0x2714))
	Wait  = string(rune(0x267B))
	Run  = string(rune(0x1F3C3))
)

func Printf(format string, a ...interface{}) {
	if len(a) > 0 {
		fmt.Printf(format+"\n", a...)
	} else {
		fmt.Println(format)
	}
}

func Fatalf(format string, a ...interface{}) {
	if len(a) > 0 {
		fmt.Printf(format+"\n", a...)
	} else {
		fmt.Println(format + "\n")
	}
	os.Exit(1)
}