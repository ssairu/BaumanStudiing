import serial
import cv2
import os
import numpy as np
import time

MAX_SIZE_FRAG = 100


class Converter:
    def __init__(self, size=4):
        self.size = size

    def int_to_bytes(self, num):
        res = b''
        for i in range(self.size):
            ost = num % 255 + 1
            res = ost.to_bytes(1, "big") + res
            num = num // 255
        return res

    def bytes_to_int(self, bint):
        mem = memoryview(bint)
        res = 0
        power = 1
        for i in range(self.size):
            res += (mem[self.size - i - 1] - 1) * power
            power *= 255
        return res


class frag:
    def __init__(self, img_id, frag_id, data):
        self.img_id = img_id
        self.id = frag_id
        self.size = len(data)
        self.data = data

    def print(self):
        print('[ img_id = ')
        print(self.img_id)
        print('\n  id = ')
        print(self.id)
        print('\n  size = ')
        print(self.size)
        print('\n  data = ')
        print(self.data)
        print(' ]\n')


idimg = 1


class image:

    def __init__(self, data, shape, max_size_frag=MAX_SIZE_FRAG):
        global idimg
        self.id = idimg
        self.max_size_frag = max_size_frag
        self.shape = shape
        idimg += 1
        if len(data) % self.max_size_frag == 0:
            self.num_frags = len(data) // self.max_size_frag
        else:
            self.num_frags = len(data) // self.max_size_frag + 1

        frags = []
        for i in range(self.num_frags):
            sizef = self.max_size_frag
            if i == self.num_frags - 1:
                sizef = len(data) - i * self.max_size_frag
            fragx = frag(self.id, i, data[i * self.max_size_frag: i * self.max_size_frag + sizef])
            frags += [fragx]
        self.frags = frags

    def tobytesFrags(self):
        bfrags = []
        c = Converter()
        lc = Converter(2)
        for x in self.frags:
            x.print()
            fragx = (b'&$&' + c.int_to_bytes(x.img_id) + c.int_to_bytes(x.id) +
                     c.int_to_bytes(x.size) + c.int_to_bytes(self.num_frags) +
                     lc.int_to_bytes(self.shape[0]) + lc.int_to_bytes(self.shape[1]) + x.data + b'%@%')
            bfrags += [fragx]
        return bfrags


# Set the file path for the source image
path = r'ps.png'

# Load the image using OpenCV
img = cv2.imread(path)
print(img.shape)

ser = serial.Serial('/dev/ttyUSB9')
ser.baudrate = 115200

devided_image = image(img.tobytes(), img.shape).tobytesFrags()
counter = 0
for x in devided_image:
    print(len(x))
    print(counter)
    counter = counter + 1
    print(x.replace(b'\x00', b'\x01'))

    ser.write(x.replace(b'\x00', b'\x01'))
    # ser.write(b'&$&------------------------------------------------------------' + counter + b'%@%')
    # counter += b'+'
    time.sleep(0.5)
    # ser.write(x)

# Закрываем порт
ser.close()
