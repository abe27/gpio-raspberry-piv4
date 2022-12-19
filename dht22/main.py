#!/usr/local/bin/python3
# -*- coding: utf-8 -*-
import DHT22 as DHT22
import time


#GPIO 21
out = 11

#c = celsius
#f = Fahrenheit
#k = kelvin
units = 'c'
sensor = DHT22.DHT22(out,units)
time.sleep(1)
running = True
while(running):
    try:
        humidity, temperature = sensor.get_value()
        if humidity is not None and temperature is not None:
            print ('Temp={0:0.1f}*C  Humidity={1:0.1f}%'.format(temperature, humidity))
        time.sleep(2)
    except KeyboardInterrupt:
        running = False