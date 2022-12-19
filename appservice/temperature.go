package main

import (
	"fmt"
	"log"

	"github.com/d2r2/go-dht"
)

func main() {
	temperature, humidity, retried, err :=
		dht.ReadDHTxxWithRetry(dht.DHT22, 11, false, 10)
	if err != nil {
		log.Fatal(err)
	}
	// Print temperature and humidity
	fmt.Printf("Pin 11 Temperature = %v*C, Humidity = %v%% (retried %d times)\n",
		temperature, humidity, retried)

	temperature, humidity, retried, err =
		dht.ReadDHTxxWithRetry(dht.DHT22, 17, false, 10)
	if err != nil {
		log.Fatal(err)
	}
	// Print temperature and humidity
	fmt.Printf("Pin 17 Temperature = %v*C, Humidity = %v%% (retried %d times)\n",
		temperature, humidity, retried)
}
