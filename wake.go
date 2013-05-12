package main

import (
	"encoding/json"
	"errors"
	"flag"
	"github.com/ghthor/gowol"
	"os"
)

func main() {
	c, err := initialize()
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		return
	}
	for _, m := range c.Macs {
		if err := wol.SendMagicPacket(m, c.Broadcast); err != nil {
			os.Stderr.WriteString("Error for MAC '" + m + "': '" + err.Error() + "'\n")
		} else {
			os.Stdout.WriteString("Waking '" + m + "' ...\n")
		}
	}
}

type config struct {
	Broadcast string
	Macs      []string
	Profiles  map[string][]string
}

func initialize() (*config, error) {
	var c config
	if err := c.load("wake.conf"); err != nil {
		os.Stderr.WriteString("Failed to load config file 'wake.conf': " + err.Error() + "\n")
	} else {
		os.Stderr.WriteString("Config file loaded.\n")
	}
	if err := c.parseFlags(); err != nil {
		os.Stderr.WriteString("Failed to parse flags: " + err.Error() + "\n")
	}
	return &c, nil
}

func (c *config) load(fileName string) error {
	fileReader, err := os.Open(fileName)
	if err != nil {
		os.Stderr.WriteString("No config file not found or read access is not allowed.\n")
		return err
	}
	dec := json.NewDecoder(fileReader)
	if err := dec.Decode(c); err != nil {
		os.Stderr.WriteString("An error occurred while reading the config file: " + err.Error() + "\n")
		return err
	}
	return nil
}

func (c *config) parseFlags() error {
	if len(os.Args) <= 1 {
		return errors.New("Usage: " + os.Args[0] + " -b <broadcast-address> <mac-address> [mac-address ...]\n")
	}
	var bcast = flag.String("b", "", "The network's broadcast address.")

	flag.Parse()

	if len(*bcast) <= 0 {
		if len(c.Broadcast) <= 0 {
			return errors.New("Please specify the network's broadcast address using the '-b' flag.")
		}
	} else {
		c.Broadcast = *bcast
	}
	var args = flag.Args()
	if len(args) <= -1 {
		return errors.New("Please specify at least one MAC-address.\n")
	}
	c.Macs = args
	return nil
}
