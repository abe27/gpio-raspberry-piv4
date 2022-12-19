#!/usr/bin/python
import Adafruit_DHT as DHT
#class for read dht22
class DHT22(object):

    def __init__(self,out=None, units = "c"):
        self.sensor = None
        self.units = units
        self.out = out
        self.sensor = DHT.DHT22

    def get_value(self):
        humidity, temp = self.read()
        temperature = getattr(self, "pass_" + self.units)(temp)
        return humidity, temperature

    def read(self):
        return DHT.read_retry(self.sensor, self.out)

    def pass_c(self, celsius):
        return celsius

    def pass_k(self, celsius):
        return celsius + 273.15

    def pass_f(self, celsius):
        return celsius * 9.0/5.0 + 32