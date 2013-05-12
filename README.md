wake
====

Tiny program for sending out wake-on-lan packets

Usage:
> ./wake -b &lt;broadcast-address&gt; &lt;mac-address&gt; \[mac-address ...\]

broadcast-address: The network's broadcast address on which to send the WOL packets.
mac-address: The MAC address of the ethernet adapter of the computer you want to wake.

For this to work, the wake-on-lan setting has to be enabled in the bios of the target computer.

TODO
----

### In progress ###
* Config-file for storing broadcast address.
* Profiles for waking 1 or more ethernet addresses by (profile) name.

### To do ###
* Describe the format of 'wake.conf'.
* Support multiple broadcast addresses, for waking devices on multiple networks simultaneously.
