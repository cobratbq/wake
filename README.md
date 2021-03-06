wake [![GoDoc](https://godoc.org/github.com/cobratbq/wake?status.svg)](https://godoc.org/github.com/cobratbq/wake) [![Build Status](https://travis-ci.org/cobratbq/wake.svg?branch=master)](https://travis-ci.org/cobratbq/wake)
====

Tiny program for sending out wake-on-lan packets

Usage:
~~~
./wake [-b <broadcast-address>] [-p profile-name] [-v] [mac-address ...]
~~~

**broadcast-address**: The network's broadcast address (ip address) on which to send the WOL packets.

**mac-address**: The MAC address of the ethernet adapter of the computer you want to wake.

**profile-name**: Profile as configured in the config file.

Either at least one mac address should be specified, or the name of a profile that is defined in the configuration file.

For this to work, the wake-on-lan setting has to be enabled in the bios of the target computer.


Example config-file
-------------------
An example config file is JSON-formatted. It must be named ' *wake.conf* '.

*wake.conf*:
~~~
{
	"broadcast": "192.168.0.255",
	"profiles": {
		"all": ["aa:aa:aa:aa:aa:aa", "bb:bb:bb:bb:bb:bb"]
	},
	"verbose": false
}
~~~
