#!/usr/local/bin/python3
# -*- coding: utf-8 -*-
import json
import requests
import sys
import DHT22 as DHT22
from random import *

api_url = "http://localhost:4040/api/v1"


def notification(token, message):
    url = "https://notify-api.line.me/api/notify"
    payload = f'message={message}'
    headers = {
        'Authorization': f'Bearer {token}',
        'Content-Type': 'application/x-www-form-urlencoded'
    }

    response = requests.request("POST", url, headers=headers, data=payload)

    print(response.text)


def get_device():
    response = requests.request("GET", f"{api_url}/notification")
    obj = response.json()
    return obj["data"]


if __name__ == '__main__':
    try:
        device = get_device()
        for i in device:
            is_accept = bool(i["is_accept"])
            serveName = i["device"]["name"]
            onPin = float(i["device"]["on_pin"])
            alertOn = float(i["device"]["alert_on"])
            print(f"serve: {serveName} pin: {onPin} alert: {alertOn}")
            # # Get temperature test
            # temperature = randint(1, 100)
            # humidity = randint(1, 100)
            # # #GPIO 17, Pin Physical 11
            # # pin = 17

            # # #c = celsius
            # # #f = Fahrenheit
            # # #k = kelvin
            units = 'c'
            sensor = DHT22.DHT22(onPin, units)
            humidity, temperature = sensor.get_value()
            if humidity is not None and temperature is not None:
                print(
                    'Temp={0:0.1f}*C  Humidity={1:0.1f}%'.format(temperature, humidity))

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
            response = requests.request(
                "POST", f"{api_url}/temp", headers=headers, data=payload)
            print(response.json())

            if temperature > alertOn:
                # notification
                if is_accept is False:
                    print(
                        f"notification: {i['line_token']['token']} is {is_accept}")
                    message = f"""\n{serveName} template: {temperature} humidity {humidity}"""
                    notification(i['line_token']['token'], message)

    except Exception as ex:
        print(ex)
        pass

    sys.exit(0)
