import serial
import cv2
import os
import numpy as np
import time
import threading

MAX_SIZE_FRAG = 3000
idimg = 1000
directory = r'/home/user/Arduino/aim'
PORT = '/dev/ttyUSB3'
commands = []
GETTING = False


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

    def getimgdata(self, not_ready=False):
        if not self.isready():
            if not_ready:
                self.sort_frags()
                data = b''
                k = 0
                for i in range(self.num_frags):
                    if k < len(self.frags) and self.frags[k].id == i:
                        data += self.frags[k].data
                        k += 1
                    else:
                        data += b'\1' * self.frags[0].size
                data = data[: self.shape[0] * self.shape[1] * 3]
                return data
            else:
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

    def add_frag(self, frag, num_fragments, img_shape):
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


def get_imges():
    print("-----------------------------------")
    global builder
    global GETTING
    ser = serial.Serial(PORT)
    # ser = serial.Serial('/dev/ttyUSB1') #WINDOWS
    ser.baudrate = 115200
    os.chdir(directory)
    with lock:
        getting = GETTING

    while getting:
        start_flag = 0
        start_aim = '&$&'
        while start_flag < 3:
            if ser.inWaiting() > 0:
                print('.')
                x = ser.read()
                # print(x)
                if str(x, 'utf-8') == start_aim[start_flag]:
                    start_flag += 1
                else:
                    start_flag *= 0
            with lock:
                if not GETTING:
                    getting = GETTING
                    print('FAlse')
                    break

        if getting:
            c = Converter()
            lc = Converter(2)
            id_img = c.bytes_to_int(ser.read(4))
            print(id_img)
            id_fragment = c.bytes_to_int(ser.read(4))
            size_fragment = c.bytes_to_int(ser.read(4))
            num_fragments = c.bytes_to_int(ser.read(4))
            shape = tuple([lc.bytes_to_int(ser.read(2)), lc.bytes_to_int(ser.read(2)), 3])
            response = ser.read(size=size_fragment)

            ending = ser.read(3)
            end_aim = b'%@%'
            if ending == end_aim:
                print('RIGHT, GET ALL')
            else:
                print('WRONG, WITHOUT ENDING')
                print("id_fragment: " + str(id_fragment))
                print("size_fragment: " + str(size_fragment))
                print("id_img: " + str(id_img))
                print("num_fragments: " + str(num_fragments))
                print("len(response): " + str(len(response)))
                print(b'ending: ' + ending)
                if b'%@%' in end_aim:
                    print('consist "%@%": YES')
                else:
                    print('consist "%@%": NO')
                continue

            frag = Frag(id_img, id_fragment, response)

            with lock:
                builder.add_frag(frag, num_fragments, shape)
                if builder.have_new():
                    for img in builder.get_new():
                        decoded_response = np.frombuffer(img.getimgdata(), dtype=np.uint8)
                        res = decoded_response.reshape(img.shape)
                        # print(decoded_response)
                        print(res)
                        filename = str(img.img_id) + 'copy.jpg'
                        cv2.imwrite(filename, res)
                    builder.clear_new()

                getting = GETTING

    ser.close()


def get_commands():
    while True:
        command = str(input())
        print("\nget command: '" + command + "'\n")
        with lock:  # Блокируем доступ к общему ресурсу commands
            commands.append(command)


def worker(command):
    global GETTING
    global thread_main
    global builder
    if command == 'stop':
        with lock:
            GETTING = False
        thread_main[0].join()
        thread_main = []
        print('\n\n reading stopped \n\n')

    if command == 'start':
        with lock:
            GETTING = True
        thread_get_imgs = threading.Thread(target=get_imges)
        thread_main.append(thread_get_imgs)
        thread_main[0].start()
    if command == 'getnotready':
        with lock:
            imgs_inload = builder.get_inload()

        for img in imgs_inload:
            decoded_response = np.frombuffer(img.getimgdata(not_ready=True), dtype=np.uint8)
            res = decoded_response.reshape(img.shape)
            print(res)
            filename = str(img.img_id) + 'copy.jpg'
            cv2.imwrite(filename, res)


def do_commands():
    global commands
    while True:
        time.sleep(0.5)
        with lock:
            if len(commands) == 0:
                continue
            else:
                for command in commands:
                    threadc = threading.Thread(target=worker, args=(command,))
                    threadc.start()
                commands = []






builder = ImageBuilder()
lock = threading.Lock()

threads = []
thread_main = []
thread_do_commands = threading.Thread(target=do_commands)
thread_do_commands.start()

thread_get_commands = threading.Thread(target=get_commands)
thread_get_commands.start()



threads.append(thread_do_commands)
threads.append(thread_get_commands)



# Ожидание завершения всех потоков
for thread in threads:
    thread.join()


