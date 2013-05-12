wake
====

Tiny program for sending out wake-on-lan packets

Usage:
> ./wake -b <broadcast-address> <mac-address> [mac-address ...]

broadcast-address: The network's broadcast address on which to send the WOL packets.
mac-address: The MAC address of the ethernet adapter of the computer you want to wake.

For this to work, the wake-on-lan setting has to be enabled in the bios of the target computer.
