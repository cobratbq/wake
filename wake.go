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
	//Initialize config first by config file.
	if err := c.load("wake.conf"); err != nil {
		os.Stderr.WriteString("Failed to load config file 'wake.conf': " + err.Error() + "\n")
	} else if c.Verbose {
		os.Stderr.WriteString("Config file loaded.\n")
	}
	//Then incorporate provided flags.
	if err := c.parseFlags(); err != nil {
		return nil, err
	}
	return &c, nil
}

func (c *config) load(fileName string) error {
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

func (c *config) parseFlags() error {
	// Define available flags.
	var bcast = flag.String("b", "", "The network's broadcast address.")
	var prof = flag.String("p", "", "The profile name of the profile to use.")
	var verbose = flag.Bool("v", false, "Be verbose during operation.")
	// Parse the command line flags
	flag.Parse()
	if len(*bcast) > 0 {
		c.Broadcast = *bcast
	} else if len(c.Broadcast) <= 0 {
		return errors.New("Please specify the network's broadcast address using the '-b' flag.")
	}
	if len(*prof) > 0 {
		addrs, ok := c.Profiles[*prof]
		if ok {
			for _, a := range addrs {
				c.Add(a)
			}
		} else {
			return errors.New("Profile with name '" + *prof + "' does not exist.")
		}
	}
	if *verbose {
		//For now only pick up the flag if it is set to true.
		c.Verbose = *verbose
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
