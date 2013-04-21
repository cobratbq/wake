package main

import (
	"errors"
	"flag"
	"github.com/ghthor/gowol"
	"os"
)

func main() {
	var c *config
	c, err := initialize()
	if err != nil {
		os.Stderr.WriteString("An error occurred while initializing: "+err.Error()+"\n")
		return
	}
	for _, m := range c.macs {
		if err := wol.SendMagicPacket(m, *c.bcast); err != nil {
			os.Stderr.WriteString("Error for MAC '" + m + "': '" + err.Error() + "'\n")
		}
	}
}

type config struct {
	bcast *string
	macs []string
}

func initialize() (*config, error) {
	var c config
	if len(os.Args) <= 1 {
		return nil, errors.New("Usage: "+os.Args[0]+" -b <broadcast-address> <mac-address> [mac-address ...]\n")
	}
	var bcast = flag.String("b", "", "The network's broadcast address.")
	flag.Parse()
	if len(*bcast) <= 0 {
		return nil, errors.New("Please specify the network's broadcast address using the '-b' flag.")
	}
	c.bcast = bcast
	var args = flag.Args()
	if len(args) <= 0 {
		return nil, errors.New("Please specify at least one MAC-address.\n")
	}
	c.macs = args
	return &c, nil
}
