package main

import (
	"flag"
	"github.com/ghthor/gowol"
	"log"
)

func main() {
	var bcast = flag.String("b", "", "The network's broadcast address.")
	flag.Parse()
	if len(*bcast) <= 0 {
		log.Printf("Please specify the network's broadcast address using the '-b' flag.")
		return
	}
	var args = flag.Args()
	if len(args) <= 0 {
		log.Printf("Please specify at least one MAC-address.\n")
		return
	}
	log.Printf("Using broadcast address: %s\n", *bcast)
	for _, m := range args {
		if err := wol.SendMagicPacket(m, *bcast); err != nil {
			log.Printf("Error for MAC '%s': '%s'\n", m, err.Error())
		}
	}
}
