#!/usr/local/bin/python3
# -*- coding: utf-8 -*-
# import DHT22 as DHT22


# #GPIO 17, Pin Physical 11
# pin = 17

# #c = celsius
# #f = Fahrenheit
# #k = kelvin
# units = 'c'
# sensor = DHT22.DHT22(pin,units)
# humidity, temperature = sensor.get_value()
# if humidity is not None and temperature is not None:
#     print ('Temp={0:0.1f}*C  Humidity={1:0.1f}%'.format(temperature, humidity))

import json
import requests
import sys
from random import *

api_url = "http://localhost:4040/api/v1"


def get_device():
    response = requests.request("GET", f"{api_url}/device")
    obj = response.json()
    return obj["data"]

if __name__ == '__main__':
    try:
        device = get_device()
        for i in device:
            serveName = i["name"]
            onPin = float(i["on_pin"])
            alertOn = float(i["alert_on"])
            print(f"serve: {serveName} pin: {onPin} alert: {alertOn}")
            ### Get temperature test
            temperature = randint(1, 100)
            humidity = randint(1, 100)

            payload = json.dumps({
                "device_id": serveName,
                "temp": temperature,
                "humidity": humidity,
                "description": "-",
                "is_active": True
            })
            headers = {
                'Content-Type': 'application/json'
            }
            response = requests.request("POST", f"{api_url}/temp", headers=headers, data=payload)
            print(response.json())

    except Exception as ex:
        print(ex)
        pass


    sys.exit(0)
