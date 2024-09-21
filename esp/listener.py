import serial
import cv2
import os
import numpy as np

MAX_SIZE_FRAG = 10000
idimg = 1000


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


class Frag:
    def __init__(self, img_id, frag_id, data):
        self.img_id = img_id
        self.id = frag_id
        self.size = len(data)
        self.data = data


class Image:
    def __init__(self, data, max_size_frag=MAX_SIZE_FRAG):
        global idimg
        self.id = idimg
        self.max_size_frag = max_size_frag
        idimg += 1
        self.num_frags = len(data) // self.max_size_frag + 1

        frags = []
        for i in range(self.num_frags):
            sizef = self.max_size_frag
            if i == self.num_frags - 1:
                sizef = len(data) - i * self.max_size_frag
            fragx = Frag(self.id, i, data[i * self.max_size_frag: i * self.max_size_frag + sizef])
            frags += [fragx]
        self.frags = frags

    def tobytesfrags(self):
        bfrags = []
        for x in self.frags:
            fragx = bytes('&$&', 'utf-8') + bytes(x.img_id) + bytes(x.id) + bytes(x.size) + bytes(
                self.num_frags) + x.data + bytes('%@%', 'utf-8')  # TODO TYPES
            bfrags += [fragx]
        return bfrags


class Imgforbuilder:
    def __init__(self, img_id, num_frags, shape):
        self.img_id = img_id
        self.frags = []
        self.num_frags = num_frags
        self.ready = False
        self.shape = shape

    def add_frag(self, frag):
        if self.ready:
            print('cannot add frag because all frags have been got')
            return False
        for x in self.frags:
            if frag.id == x.id:
                print('the same frag tried to add')
                return False
        self.frags += [frag]
        if self.isready():
            self.ready = True
        return True

    def isready(self):
        return len(self.frags) == self.num_frags

    def sort_frags(self):
        N = len(self.frags)
        for i in range(N - 1):
            for j in range(N - 1 - i):
                if self.frags[j].id > self.frags[j + 1].id:
                    self.frags[j], self.frags[j + 1] = self.frags[j + 1], self.frags[j]

    def getimgdata(self):
        if not self.isready():
            print('cannot get img because it is not full')
            return
        else:
            self.sort_frags()
            data = b''
            for f in self.frags:
                data += f.data
            return data


class ImageBuilder:
    def __init__(self):
        self.get = []
        self.load = []
        self.new = []

    def imginload(self, frag):
        for x in self.load:
            if frag.img_id == x.img_id:
                return True
        return False

    def get_img_pos_load(self, frag):
        res = -1
        for i in range(len(self.load)):
            if self.load[i].img_id == frag.img_id:
                res = i
        return res

    def get_img_pos_get(self, frag):
        res = -1
        for i in range(len(self.get)):
            if self.get[i].img_id == frag.img_id:
                res = i
        return res

    def add_frag(self, frag, img_shape):
        if not self.imginload(frag):
            self.load += [Imgforbuilder(frag.img_id, num_fragments, img_shape)]
            print("add image " + str(frag.img_id) + "to builder")
            print("frag added " + str(frag.id) + "\nfrags:")
            print(len(self.load[-1].frags))

        ipos = self.get_img_pos_load(frag)
        self.load[ipos].add_frag(frag)
        print(str(frag.id) + " / " + str(self.load[ipos].num_frags) + "total(" + str(len(self.load[ipos].frags)) + ")")
        if self.load[ipos].isready():
            self.get += [self.load[ipos]]
            self.new += [self.load[ipos]]
            self.load = self.load[: ipos] + self.load[ipos + 1:]
        # print(self.load)

    def get_ready(self):
        return self.get

    def get_inload(self):
        return self.load

    def get_new(self):
        return self.new

    def clear_new(self):
        self.new = []
        return

    def clear_get(self):
        self.get = []
        return

    def have_new(self):
        if not self.new:
            return False
        else:
            return True


# c = Converter()
# for i in range(10000000):
#     num = c.bytes_to_int(c.int_to_bytes(i))
#     if num != i:
#         print(c.bytes_to_int(c.int_to_bytes(i)))
#     if i % 1000000 == 0:
#         print(str(i) + "%\n")


# Set the directory for saving the image
directory = r'/home/user/Arduino/aim'
# directory = r'C:\Users\Dmitriy\Pictures' #WINDOWS


# Change the working directory to the specified directory for saving the image
os.chdir(directory)

ser = serial.Serial('/dev/ttyUSB5')
# ser = serial.Serial('/dev/ttyUSB1') #WINDOWS
ser.baudrate = 115200

get_imgs = []
builder = ImageBuilder()

while 0 == 0:

    start_flag = 0
    start_aim = '&$&'
    while start_flag < 3:
        x = ser.read()
        print(x)
        if str(x, 'utf-8') == start_aim[start_flag]:
            start_flag += 1
        else:
            start_flag *= 0
    c = Converter()
    lc = Converter(2)
    id_img = c.bytes_to_int(ser.read(4))
    print(id_img)
    id_fragment = c.bytes_to_int(ser.read(4))
    size_fragment = c.bytes_to_int(ser.read(4))
    num_fragments = c.bytes_to_int(ser.read(4))
    shape = tuple([lc.bytes_to_int(ser.read(2)), lc.bytes_to_int(ser.read(2)), 3])
    response = ser.read(size=size_fragment)

    ending = str(ser.read(3), 'utf-8')
    end_aim = '%@%'
    if ending == end_aim:
        print('RIGHT, GET ALL')
    else:
        print('WRONG, WITHOUT ENDING')
        break

    frag = Frag(id_img, id_fragment, response)
    builder.add_frag(frag, shape)
    # print(builder.get_new())

    if builder.have_new():
        for img in builder.get_new():
            decoded_response = np.frombuffer(img.getimgdata(), dtype=np.uint8)
            res = decoded_response.reshape(img.shape)
            # print(decoded_response)
            print(res)

            # Save the image with the filename "cat.jpg"
            filename = str(idimg) + 'copy.jpg'
            idimg += 1
            cv2.imwrite(filename, res)
        builder.clear_new()

# Закрываем порт
ser.close()
