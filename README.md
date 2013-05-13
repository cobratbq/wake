wake
====

Tiny program for sending out wake-on-lan packets

Usage:
> ./wake \[-b &lt;broadcast-address&gt;\] \[-p profile-name\] \[mac-address ...\]

broadcast-address: The network's broadcast address on which to send the WOL packets.
mac-address: The MAC address of the ethernet adapter of the computer you want to wake.
profile-name: Profile as configured in the config file.

Either at least one mac address should be specified, or the name of a profile that is defined in the configuration file.

For this to work, the wake-on-lan setting has to be enabled in the bios of the target computer.

TODO
----

### In progress ###

### Others ###
* Describe the format of 'wake.conf'.
* Support multiple broadcast addresses, for waking devices on multiple networks simultaneously.
