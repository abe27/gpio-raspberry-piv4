package main

import (
	"fmt"

	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	pin := rpio.Pin(4)

	pin.Output() // Output mode
	pin.High()   // Set pin High
	pin.Low()    // Set pin Low
	pin.Toggle() // Toggle pin (Low -> High -> Low)

	pin.Input()       // Input mode
	res := pin.Read() // Read state from pin (High / Low)
	fmt.Println(res)

	pin.Mode(rpio.Output) // Alternative syntax
	pin.Write(rpio.High)  // Alternative syntax
	rpio.Close()
}
