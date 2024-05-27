import numpy as np
from glumpy import app, gl, glm, gloo
from PIL import Image
import math

vertex = """
uniform mat4   model;         // Model matrix
uniform mat4   view;          // View matrix
uniform mat4   projection;    // Projection matrix

attribute vec3 position;      // Vertex position
attribute vec2 texture;   // Vertex texture coordinates
attribute vec3 normal;        // Vertex normal

varying vec3   v_position;      // Interpolated position (out)
varying vec2   v_texcoord;   // Interpolated fragment texture coordinates (out)
varying vec3   v_normal;        // Interpolated normal (out)
void main()
{
    // Assign varying variables
    v_normal   = normal;
    v_texcoord  = texture;
    v_position = position;
    
    gl_Position = projection * view * model * vec4(position,1.0);
}
"""

fragment = """
uniform mat4   model;         // Model matrix
uniform mat4   view;          // View matrix
uniform mat4   u_normal;          // Normal matrix
uniform mat4   projection;    // Projection matrix
uniform vec3      u_light_position;  // Light position
uniform vec3      u_light_intensity; // Light intensity
uniform sampler2D texture; // Texture


varying vec3   v_position;      // Interpolated position (out)
varying vec2   v_texcoord;   // Interpolated fragment texture coordinates (out)
varying vec3   v_normal;        // Interpolated normal (out)

void main()
{    
    // Calculate normal in world coordinates
    vec3 normal1 = normalize(u_normal * vec4(v_normal,1.0)).xyz;

    // Calculate the location of this fragment (pixel) in world coordinates
    vec3 position1 = vec3(view * model * vec4(v_position, 1));

    // Calculate the vector from this pixels surface to the light source
    vec3 surfaceToLight = u_light_position - position1;

    // Calculate the cosine of the angle of incidence (brightness)
    float brightness = dot(normal1, surfaceToLight) /
                      (length(surfaceToLight) * length(normal1));
    brightness = max(min(brightness,1.0),0.0);

    // Calculate final color of the pixel, based on:
    // 1. The angle of incidence: brightness
    // 2. The color/intensities of the light: light.intensities
    // 3. The texture and texture coord: texture(tex, fragTexCoord)

    // Get texture color
    vec4 t_color = vec4(texture2D(texture, v_texcoord).rgb, 1.0);

    // Final color
    gl_FragColor = t_color * (0.1 + 0.9*brightness * vec4(u_light_intensity, 1));
}
"""

framerate = 60
eps = 0.00001
t = 0
n = 4
r = 1
dx, dy, dz = 0, 2, 0
phi, theta, size = 0, 270, 0.8
jumping = True
v0 = 0
window = app.Window(width=512, height=512, color=(0, 0, 0, 1))


@window.event
def on_key_press(symbol, modifiers):
    global phi, theta, size, jumping
    print(symbol)
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
    elif symbol == 32:
        jumping = not jumping
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
    view = cube['view'].reshape(4, 4)
    glm.rotate(model, theta, 1, 0, 0)
    glm.rotate(model, phi, 0, 1, 0)
    glm.scale(model, size, size, size)
    if jumping:
        Jumping()
    glm.translate(model, dx, dy, dz)
    cube['model'] = model
    cube['u_normal'] = np.array(np.matrix(np.dot(view, model)).I.T)


def Jumping():
    global v0, dy, framerate, t
    if dy + 1.3 <= eps:
        v0 = -v0
    tx = 1 / framerate / 10
    t += 1
    v0 = v0 + 0.5 * tx
    dy -= v0


@window.event
def on_resize(width, height):
    cube['projection'] = glm.perspective(45.0, width / float(height), 0.1, 100.0)


@window.event
def on_init():
    gl.glEnable(gl.GL_DEPTH_TEST)
    gl.glDisable(gl.GL_BLEND)
    # gl.glEnable(gl.GL_CULL_FACE)
    # gl.glCullFace(gl.GL_FRONT)
    # gl.glFrontFace(gl.GL_CW)


V = np.zeros(n * 5, [("position", np.float32, 3),
                     ('normal', np.float32, 3),
                     ("texture", np.float32, 2)])

for i in range(n):
    for x in range(0, 4):
        V["position"][5 * i + x][0] = r * math.cos(2 * (5 * (i + x // 2)) * math.pi / (5 * n))
        V["position"][5 * i + x][1] = r * math.sin(2 * (5 * (i + x // 2)) * math.pi / (5 * n))
        V["position"][5 * i + x][2] = (x % 2) - 1

        V["texture"][5 * i + x][0] = x // 2
        V["texture"][5 * i + x][1] = (x % 2) * 0.8

        V["normal"][5 * i + x][0] = math.cos(2 * (i + 0.5) * math.pi / n)
        V["normal"][5 * i + x][1] = math.sin(2 * (i + 0.5) * math.pi / n)
        V["normal"][5 * i + x][2] = 0

    V["position"][5 * i + 4] = [0, 0, 1]
    V["texture"][5 * i + 4] = [0.5, 1]
    V["normal"][5 * i + 4][0] = math.cos(2 * (i + 0.5) * math.pi / n)
    V["normal"][5 * i + 4][1] = math.sin(2 * (i + 0.5) * math.pi / n)
    V["normal"][5 * i + 4][2] = 0.5

#
# V["position"] = [[1, 1, 1], [-1, 1, 1], [-1, -1, 1], [1, -1, 1],
#                  [1, -1, -1], [1, 1, -1], [-1, 1, -1], [-1, -1, -1]]
# V["color"] = [[0, 1, 1, 1], [0, 0, 1, 1], [0, 0, 0, 1], [0, 1, 0, 1],
#               [1, 1, 0, 1], [1, 1, 1, 1], [1, 0, 1, 1], [1, 0, 0, 1]]
V = V.view(gloo.VertexBuffer)

a = []
for i in range(n):
    first_low = i * 5
    first_up = i * 5 + 1
    second_low = i * 5 + 2
    second_up = i * 5 + 3
    upper = i * 5 + 4
    a += [first_low, second_low, first_up]
    a += [second_low, second_up, first_up]
    a += [first_up, second_up, upper]
    # a += [first_low, second_low, 0]

I = np.array(a, dtype=np.uint32)
I = I.view(gloo.IndexBuffer)

cube = gloo.Program(vertex, fragment)
cube.bind(V)

cube["u_light_position"] = -2, -2, 2
cube["u_light_intensity"] = 1, 1, 1
cube['texture'] = np.array(Image.open("./flowers.bmp"))
cube['model'] = np.eye(4, dtype=np.float32)
cube['view'] = glm.translation(0, 0, -5)
cube['projection'] = glm.perspective(45.0, 1, 0.1, 100.0)

app.run(framerate=framerate)
