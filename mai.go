package main

import (
	"MinerClient/pkg/X19"
	"fmt"
)

func main() {
	bb := X19.NewDevice()
	bb.Ip.Host = "10.29.4.5"
	stats, err := bb.GetStats()
	if err != nil {
		return
	}
	fmt.Println(stats)
}
