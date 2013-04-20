package main

import (
	"flag"
	"github.com/ghthor/gowol"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		os.Stderr.WriteString("Usage: "+os.Args[0]+" -b <broadcast-address> <mac-address> [mac-address ...]\n")
		return
	}
	var bcast = flag.String("b", "", "The network's broadcast address.")
	flag.Parse()
	if len(*bcast) <= 0 {
		os.Stderr.WriteString("Please specify the network's broadcast address using the '-b' flag.")
		return
	}
	var args = flag.Args()
	if len(args) <= 0 {
		os.Stderr.WriteString("Please specify at least one MAC-address.\n")
		return
	}
	for _, m := range args {
		if err := wol.SendMagicPacket(m, *bcast); err != nil {
			os.Stderr.WriteString("Error for MAC '" + m + "': '" + err.Error() + "'\n")
		}
	}
}
