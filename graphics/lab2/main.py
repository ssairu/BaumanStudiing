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


def window_resize(window, width, height):
    glViewport(0, 0, width, height)


vertex_code = """
    # version 330

    layout(location = 0) attribute vec3 position;
    layout(location = 1) attribute vec3 color;
    varying vec4 v_color;

    void main(){
        gl_Position = vec4(position, 1.0);
        v_color = vec4(color, 1.0);
    } """

fragment_code = """
    # version 330

    varying vec4 v_color;

    void main()
    {
      gl_FragColor = v_color;
    } """

glfw.terminate()

glfw.init()
window = glfw.create_window(900, 700, "PyOpenGL lab1", None, None)
glfw.set_window_pos(window, 400, 200)
glfw.make_context_current(window)
glfw.set_key_callback(window, key_callback)
glfw.set_mouse_button_callback(window, mouse_button_callback)
glfw.set_window_size_callback(window, window_resize)

vertices = [(-0.9, -0.5, 0.0),
            (0.0, -0.9, 0.0),
            (0.5, -0.5, 0.0),
            (0.5, 0.5, 0.0),
            (-0.5, 0.5, 0.0)]

# list the color code here
colors = [(1, 0.5, 0),
          (0, 0.8, 0.9),
          (0, 0.3, 0.6),
          (0.1, 0.8, 0.1),
          (1, 0.2, 1.0)]

indexes = [0, 1, 2,
           0, 2, 3,
           0, 3, 4]

program = glCreateProgram()
vertex = glCreateShader(GL_VERTEX_SHADER)
fragment = glCreateShader(GL_FRAGMENT_SHADER)

# Set shaders source
glShaderSource(vertex, vertex_code)
glShaderSource(fragment, fragment_code)

# Compile shaders
glCompileShader(vertex)
if not glGetShaderiv(vertex, GL_COMPILE_STATUS):
    error = glGetShaderInfoLog(vertex).decode()
    print(error)
    raise RuntimeError("Vertex shader compilation error")

glCompileShader(fragment)
if not glGetShaderiv(fragment, GL_COMPILE_STATUS):
    error = glGetShaderInfoLog(fragment).decode()
    print(error)
    raise RuntimeError("Fragment shader compilation error")

glAttachShader(program, vertex)
glAttachShader(program, fragment)
glLinkProgram(program)

if not glGetProgramiv(program, GL_LINK_STATUS):
    print(glGetProgramInfoLog(program))
    raise RuntimeError('Linking error')

glDetachShader(program, vertex)
glDetachShader(program, fragment)
glUseProgram(program)

buffer = glGenBuffers(1)
glBindBuffer(GL_ARRAY_BUFFER, buffer)

indexes = np.array(indexes, np.uint32)
data = np.zeros(5, [("position", np.float32, 3),
                    ("color", np.float32, 3)])
data['position'] = vertices
data['color'] = colors

stride = data.strides[0]
offset = ctypes.c_void_p(0)
glEnableVertexAttribArray(0)
glBindBuffer(GL_ARRAY_BUFFER, buffer)
glVertexAttribPointer(0, 3, GL_FLOAT, False, stride, offset)

offset = ctypes.c_void_p(data.dtype["position"].itemsize)
glEnableVertexAttribArray(1)
glBindBuffer(GL_ARRAY_BUFFER, buffer)
glVertexAttribPointer(1, 3, GL_FLOAT, False, stride, offset)
glBufferData(GL_ARRAY_BUFFER, data.nbytes, data, GL_DYNAMIC_DRAW)

EBO = glGenBuffers(1)
glBindBuffer(GL_ELEMENT_ARRAY_BUFFER, EBO)
glBufferData(GL_ELEMENT_ARRAY_BUFFER, indexes.nbytes, indexes, GL_STATIC_DRAW)

# Now we will pour color for the animation's background
glClearColor(0, 0.7, 0.5, 1)

while not glfw.window_should_close(window):
    # this while loop will keep iterating all the functions below until the window is not closed
    glfw.poll_events()
    glClear(GL_COLOR_BUFFER_BIT)
    # creating rotation animated motion
    glRotatef(angle, 1, 1, 1)
    # glDrawArrays(GL_POLYGON, 0, 5)
    glDrawElements(GL_TRIANGLES, len(indexes), GL_UNSIGNED_INT, None)
    glfw.swap_buffers(window)
