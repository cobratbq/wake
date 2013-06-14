package main

import (
	"encoding/json"
	"errors"
	"flag"
	"github.com/ghthor/gowol"
	"os"
	"strings"
)

func main() {
	c, err := initialize()
	if err != nil {
		os.Stderr.WriteString("Initialization failed: " + err.Error() + "\n")
		return
	}
	for m, _ := range c.Macs {
		if err := wol.SendMagicPacket(m, c.Broadcast); err != nil {
			os.Stderr.WriteString("Error for MAC '" + m + "': '" + err.Error() + "'\n")
		} else if c.Verbose {
			os.Stdout.WriteString("Waking '" + m + "' ...\n")
		}
	}
}

type config struct {
	Broadcast string
	Macs      map[string]bool
	Profiles  map[string][]string
	Verbose   bool
}

func initialize() (*config, error) {
	//Initialize config struct
	var c config
	c.Macs = make(map[string]bool)
	c.Profiles = make(map[string][]string)
	//Initialize flags
	var flgs = initFlags()
	//Initialize config first by config file.
	if err := c.loadConfig("wake.conf"); err != nil && *flgs.Verbose {
		os.Stderr.WriteString("Failed to load config file 'wake.conf': " + err.Error() + "\n")
	} else if c.Verbose {
		os.Stderr.WriteString("Config file loaded.\n")
	}
	//Then incorporate provided flags.
	if err := c.loadFlags(flgs); err != nil {
		return nil, err
	}
	return &c, nil
}

func (c *config) loadConfig(fileName string) error {
	fileReader, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer fileReader.Close()
	dec := json.NewDecoder(fileReader)
	if err := dec.Decode(c); err != nil {
		return err
	}
	return nil
}

type flags struct {
	Bcast   *string
	Prof    *string
	Verbose *bool
}

func initFlags() *flags {
	var f flags
	// Define available flags.
	f.Bcast = flag.String("b", "", "The network's broadcast address.")
	f.Prof = flag.String("p", "", "The profile name of the profile to use.")
	f.Verbose = flag.Bool("v", false, "Be verbose during operation.")
	flag.Parse()
	return &f
}

func (c *config) loadFlags(flgs *flags) error {
	// Parse the command line flags
	if len(*flgs.Bcast) > 0 {
		c.Broadcast = *flgs.Bcast
	} else if len(c.Broadcast) <= 0 {
		return errors.New("Please specify the network's broadcast address using the '-b' flag.")
	}
	if len(*flgs.Prof) > 0 {
		addrs, ok := c.Profiles[*flgs.Prof]
		if ok {
			for _, a := range addrs {
				c.Add(a)
			}
		} else {
			return errors.New("Profile with name '" + *flgs.Prof + "' does not exist.")
		}
	}
	if *flgs.Verbose {
		//For now only pick up the flag if it is set to true.
		c.Verbose = *flgs.Verbose
	}
	var args = flag.Args()
	if len(args) > 0 {
		for _, address := range args {
			c.Add(address)
		}
	}
	return nil
}

func (c *config) Add(macAddress string) {
	var mac = strings.ToLower(macAddress)
	c.Macs[mac] = true
}
