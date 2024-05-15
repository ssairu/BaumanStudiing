from itertools import chain

import numpy as np
from OpenGL.GL import *
import glfw

angle = 0.0


def mouse_button_callback(window, button, action, mods):
    global angle
    if action == glfw.PRESS:
        if button == glfw.MOUSE_BUTTON_RIGHT:
            angle = -0.01
        if button == glfw.MOUSE_BUTTON_LEFT:  # glfw.KEY_LEFT
            angle = 0.1
#
def key_callback(window, key, scancode, action, mods):
    global angle
    if action == glfw.PRESS:
        angle = 0.0


glfw.terminate()

glfw.init()
window = glfw.create_window(900, 700, "PyOpenGL lab1", None, None)
glfw.set_window_pos(window, 400, 200)
glfw.make_context_current(window)
glfw.set_key_callback(window, key_callback)
glfw.set_mouse_button_callback(window, mouse_button_callback)

a = [
                [10, 20], [250, 200], [300, 30], [400, 300], [300, 400]
            ]
z=[]
for x in a:
    x.append(0)
    z.append(x)

x = list(chain.from_iterable(z))
print(x)
vertices = [y/800 for y in x]
print(vertices)

# list the color code here
colors = [0.5, 0.5, 0,
          0, 0.8, 0.9,
          0, 0.3, 0.6,
          0.1, 0.8, 0.1,
          1, 0.2, 1.0]

v = np.array(vertices, dtype=np.float32)
c = np.array(colors, dtype=np.float32)
# this will create a color-less triangle
glEnableClientState(GL_VERTEX_ARRAY)
glVertexPointer(3, GL_FLOAT, 0, v)
glEnableClientState(GL_COLOR_ARRAY)
glColorPointer(3, GL_FLOAT, 0, c)
# Now we will pour color for the animation's background
glClearColor(0, 0.7, 0.5, 1)

while not glfw.window_should_close(window):
    # this while loop will keep iterating all the functions below until the window is not closed
    glfw.poll_events()
    glClear(GL_COLOR_BUFFER_BIT)
    # creating rotation animated motion
    glRotatef(angle, 1, 1, 1)
    glDrawArrays(GL_POLYGON, 0, 5)
    glfw.swap_buffers(window)
