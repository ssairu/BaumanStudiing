import serial
import cv2 
import os 
import numpy as np




class frag:
	def __init__(self, img_id, frag_id, data):
		self.img_id = img_id
		self.id = frag_id
		self.size = len(data)
		self.data = data
		
	def getfrag(bytes):
		mem = memoryview(bytes)
		starting = str(mem[:3])
		start_aim = '&$&'
		if (starting == start_aim):
			print('RIGHT, GET STARTED')
		else:
			print('WRONG, WITHOUT START')
			return
		
		
		id_img = int(mem[3:7])
		id_fragment = int(mem[7:11])
		size_fragment = int(mem[11:15])
		num_fragments = int(mem[15:19])
		response = bytes(mem[19:-3])
		
		
		ending = str(mem[-3:])
		end_aim = '%@%'
		if (ending == end_aim):
			print('RIGHT, GET ALL')
		else:
			print('WRONG, WITHOUT ENDING')
			return
			
		return frag(id_img, id_fragment, response)


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
		
	
	
class imgforbuilder:
	def __init__(self, img_id, num_frags):
		self.img_id = img_id
		self.frags = []
		self.num_frags = num_frags
		self.ready = False
		
		
	def add_frag(self, frag):
		if (self.ready):
			print('cannot add frag because all frags have been got')
			return False
		for x in self.frags:
			if (frag.id == x.id):
				print('the same frag tried to add')
				return False
		self.frags += [frag]
		if (self.isready()):
			self.ready = True
		return True
		
		
	def isready(self):
		return len(self.frags) == self.num_frags
		
		
	def getimg(self):
		if (!self.isready()):
			print('cannot get img because it is not full')
		else:
			data = b''
			for f in self.frags:
				
			
			
		
		
class image_builder:

	
	
	def __init__(self):
		self.get = []	
		self.load = []
		
	def imginload(self, frag):
		for x in self.load:
			if (frag.img_id == x.img_id):
				return True
		return False
		
	
	
	def add_frag(self, frag):
		if (self.imginload(frag)):
			




# Set the directory for saving the image
directory = r'/home/user/Arduino/aim'
# Change the working directory to the specified directory for saving the image
os.chdir(directory) 


ser = serial.Serial('/dev/ttyUSB1')
ser.baudrate = 115200

get_imgs = []

for i in range(1):


	start_flag = 0
	start_aim = '&$&'
	while(start_flag < 3):
		x = ser.read()
		print(x)
		if (x == start_aim[start_flag]):
			start_flag += 1
		else:
			start_flag *= 0
	
	
	id_img = ser.read(4)
	id_fragment = ser.read(4)
	size_fragment = ser.read(4)
	num_fragments = ser.read(4)
	response = ser.read(size=size_fragment)
	
	
	ending = ser.read(3)
	end_aim = '%@%'
	if (ending == end_aim):
		print('RIGHT, GET ALL')
	else:
		print('WRONG, WITHOUT ENDING')
		
		
		
	print(response)
	decoded_response = np.frombuffer(response, dtype=uint8)
	res = decoded_response.reshape(img.shape)
	print(decoded_response)
	
	# Save the image with the filename "cat.jpg"
	filename = 'copy.jpg'
	cv2.imwrite(filename, res)
	

# Закрываем порт
ser.close()

