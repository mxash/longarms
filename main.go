package main

import (
	"log"
	"os/exec"
	"strings"

	evdev "github.com/grafov/evdev"

	"github.com/spf13/viper"
)

// runCom gets the command string from the config file based off of the type of tap, parses it into a splice, and runs it using exec.Command
func runCom(tap string) error {
	// Get command string from the config file based off of the number of fingers tapped
	command := viper.GetString(tap)
	// Parse command string into a slice based off of whitespace
	args := strings.Fields(command)
	// Separate the command string from the args
	com := args[0]
	// Delete the command string from the slice
	args = args[1:]

	// create exec.Command instance using com and args
	cmd := exec.Command(com, args...)
	// Run exec.Command instance
	err := cmd.Run()

	log.Printf("%s tapped!", tap)

	return err
}

func main() {
	// Global variables for device and err so that they can be set within an if statement
	var device *evdev.InputDevice
	var err error

	// Set config name, file type, and add relevant paths
	viper.SetConfigName("longarms")
	viper.AddConfigPath("$HOME/.config/longarms")

	// Watch config based off the name, type, and path above
	viper.ReadInConfig()

	// Check if an evdev device is configured
	if viper.IsSet("device") {
		device, err = evdev.Open(viper.GetString("device"))
		if err != nil {
			log.Panic(err)
		}
	} else {
		log.Panic("No input device set in config.")
	}

	for {
		events, err := device.Read()
		if err != nil {
			log.Fatal(err)
		}
		// Loop through events in Evdev device
		for i := range events {
			ev := &events[i]

			if ev.Type == evdev.EV_KEY && ev.Value == 1 {
				val, hasbtn := evdev.BTN[int(ev.Code)]
				if hasbtn {
					switch val {
					case "BTN_TOOL_TRIPLETAP":
						if viper.IsSet("tripletap") {
							err := runCom("tripletap")
							if err != nil {
								log.Fatal(err)
							}
						}
					case "BTN_TOOL_QUADTAP":
						if viper.IsSet("quadtap") {
							err := runCom("quadtap")
							if err != nil {
								log.Fatal(err)
							}
						}
					case "BTN_TOOL_QUINTTAP":
						if viper.IsSet("quintap") {
							err := runCom("quintap")
							if err != nil {
								log.Fatal(err)
							}
						}
					}
				}
			}
		}
	}
}
