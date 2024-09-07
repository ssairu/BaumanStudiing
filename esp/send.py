import serial
import cv2
import os
import numpy as np

MAX_SIZE_FRAG = 100

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

    def __init__(self, data, max_size_frag=MAX_SIZE_FRAG):
        global idimg
        self.id = idimg
        self.max_size_frag = max_size_frag
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
        for x in self.frags:
            x.print()
            fragx = (b'&$&' + x.img_id.to_bytes(4, 'big') + x.id.to_bytes(4, 'big') +
                     x.size.to_bytes(4, 'big') + self.num_frags.to_bytes(4, 'big') +
                     x.data + b'%@%')
            bfrags += [fragx]
        return bfrags


# Set the file path for the source image
path = r'hen.jpg'

# Load the image using OpenCV
img = cv2.imread(path)

ser = serial.Serial('/dev/ttyUSB1')
ser.baudrate = 115200

devided_image = image(img.tobytes()).tobytesFrags()

for x in devided_image:

    print(len(x))
    print(bytes('&$&', 'utf-8'))
    print(x)
    ser.write(x)
    ser.write(b'&$&-------------------------------------------------------------%@%')

# Закрываем порт
ser.close()
