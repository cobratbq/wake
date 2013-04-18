package main

import (
	"fmt"
	"github.com/ghthor/gowol"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <broadcast-address> <mac-address> [mac-address ...]\n", os.Args[0])
		return
	}
	bcast := os.Args[1]
	macs := os.Args[2:]
	for _, m := range macs {
		if err := wol.SendMagicPacket(m, bcast); err != nil {
			fmt.Printf("Error for MAC '%s': '%s'\n", m, err.Error())
		}
	}
}
