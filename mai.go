package main

import (
	"MinerClient/pkg/X19"
	"fmt"
)

func main() {
	bb := X19.NewDevice()
	bb.Target.Host = "10.29.4.5"
	stats, err := bb.GetSystemInfo()
	fmt.Printf("%+v", stats)
	fmt.Println(err)
}
