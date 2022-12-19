package main

import (
	"fmt"

	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	fmt.Println("Start service")
	if err := rpio.Open(); err != nil {
		panic(err)
	}

	pin := rpio.Pin(4)
	pin.Output()          // Output mode
	pin.High()            // Set pin High
	pin.Low()             // Set pin Low
	pin.Toggle()          // Toggle pin (Low -> High -> Low)
	pin.Input()           // Input mode
	res := pin.Read()     // Read state from pin (High / Low)
	pin.Mode(rpio.Output) // Alternative syntax
	pin.Write(rpio.High)  // Alternative syntax
	fmt.Println(res)
	rpio.Close()
}
