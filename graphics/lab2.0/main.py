import numpy as np
from glumpy import app, gl, glm, gloo
import math

vertex = """
uniform mat4   model;         // Model matrix
uniform mat4   view;          // View matrix
uniform mat4   projection;    // Projection matrix
attribute vec4 color;         // Vertex color
attribute vec3 position;      // Vertex position
varying vec4   v_color;       // Interpolated fragment color (out)
void main()
{
    v_color = color;
    gl_Position = projection * view * model * vec4(position,1.0);
}
"""

fragment = """
varying vec4 v_color; // Interpolated fragment color (in)
void main()
{
    gl_FragColor = v_color;
}
"""

window = app.Window(width=512, height=512, color=(0, 1, 1, 1))


@window.event
def on_key_press(symbol, modifiers):
    global phi, theta, size
    # print(symbol)
    if symbol == 65362:
        theta += 1
    elif symbol == 65364:
        theta -= 1
    elif symbol == 65363:
        phi -= 1
    elif symbol == 65361:
        phi += 1
    elif symbol == 46:
        size += 0.1
    elif symbol == 44 and size > 0.2:
        size -= 0.1

    pass


@window.event
def on_draw(dt):
    window.clear()

    # Filled cube
    cube.draw(gl.GL_TRIANGLES, I)

    # Rotate cube
    model = np.eye(4, dtype=np.float32)
    glm.rotate(model, theta, 1, 0, 0)
    glm.rotate(model, phi, 0, 1, 0)
    glm.scale(model, size, size, size)
    cube['model'] = model


@window.event
def on_resize(width, height):
    cube['projection'] = glm.perspective(45.0, width / float(height), 0.1, 100.0)
    # print(glm.perspective(45.0, width / float(height), 0.1, 100.0))
    # projection = np.eye(4, dtype=np.float32)
    # glm.scale(projection, 0.75)
    # for i in range(4):
    #     projection[3][i] = 0.75
    # print(projection)
    # projection[3][1] = 0.5 * math.cos(math.pi / 4)
    # projection[3][2] = 0.5 * math.sin(math.pi / 4)
    # print(projection)
    # cube['projection'] = projection


@window.event
def on_init():
    gl.glEnable(gl.GL_DEPTH_TEST)
    # gl.glEnable(gl.GL_CULL_FACE)
    # gl.glCullFace(gl.GL_FRONT)
    # gl.glFrontFace(gl.GL_CW)


n = 56
r = 1
V = np.zeros(n * 2 + 1, [("position", np.float32, 3),
                         ("color", np.float32, 4)])

for i in range(n):
    V["position"][2 * i][0] = r * math.cos(2 * i * math.pi / n)
    V["position"][2 * i][1] = r * math.sin(2 * i * math.pi / n)
    V["position"][2 * i][2] = -1
    V["color"][2 * i] = [0.5 / float(i + 1), 1 / float(i + 1), 0.3 / float(i + 1), 1]
    V["color"][2 * i + 1] = [0.5 / float(i + 1), 1 / float(i + 1), 0.3 / float(i + 1), 1]

    V["position"][2 * i + 1][0] = r * math.cos(2 * i * math.pi / n)
    V["position"][2 * i + 1][1] = r * math.sin(2 * i * math.pi / n)
    V["position"][2 * i + 1][2] = 0
    V["color"][2 * i] = [0.5 / float(i + 1), 1 / float(i + 1), 0.3 / float(i + 1), 1]
    V["color"][2 * i + 1] = [0.5 / float(i + 1), 1 / float(i + 1), 0.3 / float(i + 1), 1]

V["position"][2 * n] = [0, 0, 1]
V["color"][2 * n] = [1, 1, 1, 1]

#
# V["position"] = [[1, 1, 1], [-1, 1, 1], [-1, -1, 1], [1, -1, 1],
#                  [1, -1, -1], [1, 1, -1], [-1, 1, -1], [-1, -1, -1]]
# V["color"] = [[0, 1, 1, 1], [0, 0, 1, 1], [0, 0, 0, 1], [0, 1, 0, 1],
#               [1, 1, 0, 1], [1, 1, 1, 1], [1, 0, 1, 1], [1, 0, 0, 1]]
V = V.view(gloo.VertexBuffer)

a = []
for i in range(n):
    next = (i + 1) % n
    first_low = i * 2
    first_up = i * 2 + 1
    second_low = next * 2
    second_up = next * 2 + 1
    a += [first_low, second_low, first_up]
    a += [second_low, second_up, first_up]
    a += [first_up, second_up, 2 * n]
    a += [first_low, second_low, 0]

I = np.array(a, dtype=np.uint32)
I = I.view(gloo.IndexBuffer)

cube = gloo.Program(vertex, fragment)
cube.bind(V)

cube['model'] = np.eye(4, dtype=np.float32)
cube['view'] = glm.translation(0, 0, -5)
# cube['projection'] = glm.perspective(45.0, 1, 0.1, 100.0)

phi, theta, size = 0, 0, 1

app.run(framerate=60)
