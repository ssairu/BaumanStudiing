import serial
import cv2 
import os 
import numpy as np




MAX_SIZE_FRAG = 10000


class frag:
	def __init__(self, img_id, frag_id, data):
		self.img_id = img_id
		self.id = frag_id
		self.size = len(data)
		self.data = data
		


class image:
	idimg = 1

	def __init__(self, data, max_size_frag = MAX_SIZE_FRAG):
		self.id = idimg
		self.max_size_frag = max_size_frag
		idimg += 1
		self.num_frags = len(data) // self.max_size_frag + 1
		
		frags = []
		for i in range(self.num_frags):
			sizef = self.max_size_frag
			if (i == self.num_frags - 1):
				sizef = len(data) - i * self.max_size_frag
			fragx = frag(self.id, i, data[i * self.max_size_frag : i * self.max_size_frag + sizef])
			frags += [fragx]
		self.frags = frags
		
	
	def tobytesFrags(self):
		bfrags = []
		for x in self.frags:
			fragx = bytes('&$&', 'utf-8') + bytes(x.img_id) + bytes(x.id) + bytes(x.size) + bytes(self.num_frags) + x.data + bytes('%@%', 'utf-8') #TODO TYPES
			bfrags += [fragx]
		return bfrags
		
	



# Set the file path for the source image
path = r'hen.jpg'

# Set the directory for saving the image
directory = r'/home/user/Arduino/aim'

# Load the image using OpenCV
img = cv2.imread(path) 

# Change the working directory to the specified directory for saving the image
os.chdir(directory) 


ser = serial.Serial('/dev/ttyUSB1')
ser.baudrate = 115200


devided_image = image(img.tobytes())

for x in devided_image.tobytesFrags():
	ser.write(x)








for i in range(1):
	#ser.write(bytes(str(i) + src, 'utf-8'))  # выводим какие-либо данные
	print(img)
	print(img.shape, img.dtype)
	ser.write()  # выводим какие-либо данные
	print(img.tobytes())w
	print('\n\n\n\n\n\n\n\n')
	

# Закрываем порт
ser.close()

