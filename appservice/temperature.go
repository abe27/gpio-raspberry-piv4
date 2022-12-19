package main

import (
	"fmt"

	"github.com/morus12/dht22"
)

func main() {
	sensor := dht22.New("GPIO_17")
	temperature, err := sensor.Temperature()
	if err != nil {
		panic(err)
	}

	humidity, err := sensor.Humidity()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Temp=%f*C  Humidity=%f%\n", temperature, humidity)
}
