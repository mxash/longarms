# Longarms

Longarms is a simple tool to recognize multitouch taps (3-finger taps, 4-finger taps, and 5-finger taps where supported) using evdev.

## Configuration

Longarms looks for config files in ```$HOME/.config/longarms```. Refer to [longarms.yaml.example](longarms.yaml.example) for options. You can run ```sudo evtest``` to find the path to your touchpad.

## Installation

### Debian/Ubuntu based distros

Installation is possible using our .deb files on our [releases page](https://github.com/mxash/longarms/releases).

### Other Distributions

There's multiple options to install with other distros, you could simply run ```go install``` if you have a Go environment setup. Otherwise you can add the binary from our tar.gz packages on our [releases page](https://github.com/mxash/longarms/releases) into a folder in your PATH.