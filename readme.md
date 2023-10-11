# Longarms

Longarms is a simple tool to recognize multitouch taps (3-finger taps, 4-finger taps, and 5-finger taps where supported) using evdev.

## Configuration

Longarms looks for config files in ```$HOME/.config/longarms```. Refer to [longarms.yaml.example](longarms.yaml.example) for options. You can run ```sudo evtest``` to find the path to your touchpad.

## Installation

Currently installation is offered simply by pulling the GitHub repo and running ```go install```. This requires your $GOPATH environment variable to be properly set up so that $GOPATH/bin is in $PATH. I may offer packages for various distros in the future, but don't currently have the tooling for this.
