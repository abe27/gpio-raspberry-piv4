#!/usr/local/bin/python3
# -*- coding: utf-8 -*-
import json
import requests
import sys
import DHT22 as DHT22
from datetime import datetime
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
            units = 'c'
            sensor = DHT22.DHT22(onPin, units)
            humidity, temperature = sensor.get_value()
            if humidity is not None and temperature is not None:
                print(
                    'Temp={0:0.1f}*C  Humidity={1:0.1f}%'.format(temperature, humidity))

            if is_accept:
                print(
                    f"notification: {i['line_token']['token']} is {is_accept}")
                d = datetime.now().strftime('%Y-%m-%d %H:%M:%S')
                message = f"""\nNow.{serveName} works normally.\nTemperature: {round(temperature, 2)}\nHumidity: {round(humidity, 2)}\nAt: {d}"""
                notification(i['line_token']['token'], message)

    except Exception as ex:
        print(ex)
        pass

    sys.exit(0)
