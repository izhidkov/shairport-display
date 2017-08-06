# shairport-display

[Shairport-sync](https://github.com/mikebrady/shairport-sync) is an open source implementation of Apple's AirPlay protocol, allowing audio streaming from iOS and macOS devices, as well as iTunes. When playing music from iTunes, metadata is sent including track information and album artwork.

**Shairport-display** is a small Go program that parses track metadata provided from shairport-sync, and displays it using OpenVG on the screen of a Raspberry Pi or similar device. By using [OpenVG bindings for Go](https://github.com/ajstarks/openvg), we can draw graphics to the Raspberry Pi's display without needing Xorg. This keeps the Raspberry Pi's system image smaller and simpler.

## This is a work-in-progress

It mostly works, but is still pretty rough. Pull requests welcome!
