from itertools import chain

import glfw
from OpenGL.GL import *

visible = False
testing = False

INSIDE = 0  # 0000
LEFT = 1  # 0001
RIGHT = 2  # 0010
BOTTOM = 4  # 0100
TOP = 8  # 1000


class Lab4:
    def __init__(self, width, height, title):
        if not glfw.init():
            return

        self.window = glfw.create_window(width, height, title, None, None)

        if not self.window:
            glfw.terminate()
            return

        self.title, self.width, self.height = title, width, height
        self.field = [0] * (width * height)
        self.cutter = {"x_max": 400, "y_max": 400, "x_min": 0, "y_min": 0}

        if testing:
            self.figure = [
                [10, 20], [250, 200], [300, 60], [400, 300], [300, 400]
            ]

        glfw.make_context_current(self.window)
        glfw.set_key_callback(self.window, self.key_callback)
        glfw.set_mouse_button_callback(self.window, self.mouseCallback)

    # Function to compute region code for a point(x, y)
    def computeCode(self, x, y):
        code = INSIDE
        if x < self.cutter["x_min"]:  # to the left of rectangle
            code |= LEFT
        elif x > self.cutter["x_max"]:  # to the right of rectangle
            code |= RIGHT
        if y < self.cutter["y_min"]:  # below the rectangle
            code |= BOTTOM
        elif y > self.cutter["y_max"]:  # above the rectangle
            code |= TOP
        return code

    def cohenSutherlandClip(self, x1, y1, x2, y2):

        # Compute region codes for P1, P2
        code1 = self.computeCode(x1, y1)
        code2 = self.computeCode(x2, y2)
        accept = False

        while True:

            # If both endpoints lie within rectangle
            if code1 == 0 and code2 == 0:
                accept = True
                break

            # If both endpoints are outside rectangle
            elif (code1 & code2) != 0:
                break

            # Some segment lies within the rectangle
            else:

                # Line needs clipping
                # At least one of the points is outside,
                # select it
                x = 1.0
                y = 1.0
                if code1 != 0:
                    code_out = code1
                else:
                    code_out = code2

                # Find intersection point
                # using formulas y = y1 + slope * (x - x1),
                # x = x1 + (1 / slope) * (y - y1)
                if code_out & TOP:
                    # Point is above the clip rectangle
                    x = x1 + (x2 - x1) * (self.cutter["y_max"] - y1) / (y2 - y1)
                    y = self.cutter["y_max"]
                elif code_out & BOTTOM:
                    # Point is below the clip rectangle
                    x = x1 + (x2 - x1) * (self.cutter["y_min"] - y1) / (y2 - y1)
                    y = self.cutter["y_min"]
                elif code_out & RIGHT:
                    # Point is to the right of the clip rectangle
                    y = y1 + (y2 - y1) * (self.cutter["x_max"] - x1) / (x2 - x1)
                    x = self.cutter["x_max"]
                elif code_out & LEFT:
                    # Point is to the left of the clip rectangle
                    y = y1 + (y2 - y1) * (self.cutter["x_min"] - x1) / (x2 - x1)
                    x = self.cutter["x_min"]

                # Now intersection point (x, y) is found
                # We replace point outside clipping rectangle
                # by intersection point
                if code_out == code1:
                    x1 = x
                    y1 = y
                    code1 = self.computeCode(x1, y1)
                else:
                    x2 = x
                    y2 = y
                    code2 = self.computeCode(x2, y2)

        if accept:
            print("Line accepted from %.2f, %.2f to %.2f, %.2f" % (x1, y1, x2, y2))
            return {"visible": True, "v_coords": [[x1, y1], [x2, y2]]}

            # Here the user can add code to display the rectangle
            # along with the accepted (portion of) lines

        else:
            return {"visible": False, "v_coords": []}

    def cutFigure(self):
        newF = []
        n = len(self.figure)
        for i in range(n):
            newline = self.cohenSutherlandClip(self.figure[i][0],
                                                  self.figure[i][1],
                                                  self.figure[(i+1)%n][0],
                                                  self.figure[(i+1)%n][1])
            if (newline["visible"]):
                if newline["v_coords"][0] != self.figure[i]:
                    newF += [newline["v_coords"][0]]
                else:
                    newF += [self.figure[i]]
                if newline["v_coords"][1] != self.figure[(i+1)%n]:
                    newF += [newline["v_coords"][1]]
                else:
                    newF += [self.figure[(i+1)%n]]

        tempF = []
        for i in range(len(newF)):
            if newF[i] == newF[(i+1)%len(newF)]:
                continue
            tempF += [newF[i]]
        newF = tempF
        self.figure = newF
        print(newF)

    def drawFigure(self, figure):
        # for i in range(len(figure)):
        #    figure[i] = [figure[i][0], self.height - figure[i][1]]

        def plot_line_low(x0, y0, x1, y1):
            dx = x1 - x0
            dy = y1 - y0
            yi = 1 if dy > 0 else -1
            dy = abs(dy)
            brightness = 1.0

            D = 2 * dy - dx
            y = y0

            for x in range(x0, x1 + 1):
                self.set_pixel(x, y, int(brightness * 255))
                if D > 0:
                    y = y + yi
                    D = D - 2 * dx
                D = D + 2 * dy

        def plot_line_high(x0, y0, x1, y1):
            dx = x1 - x0
            dy = y1 - y0
            xi = 1 if dx > 0 else -1
            dx = abs(dx)
            brightness = 1.0

            D = 2 * dx - dy
            x = x0

            for y in range(y0, y1 + 1):
                self.set_pixel(x, y, int(brightness * 255))
                if D > 0:
                    x = x + xi
                    D = D - 2 * dy
                D = D + 2 * dx

        # рисуем линии (границы)
        # с помощью алгоритма Брезенхема с устранением ступенчатости
        for i in range(0, len(figure)):
            x0 = int(figure[i][0])
            y0 = int(figure[i][1])
            x1 = int(figure[(i + 1) % len(figure)][0])
            y1 = int(figure[(i + 1) % len(figure)][1])

            if abs(y1 - y0) < abs(x1 - x0):
                if x0 > x1:
                    plot_line_low(x1, y1, x0, y0)
                else:
                    plot_line_low(x0, y0, x1, y1)
            else:
                if y0 > y1:
                    plot_line_high(x1, y1, x0, y0)
                else:
                    plot_line_high(x0, y0, x1, y1)


    def getRebra(self):
        rebra = []
        for v in range(0, len(self.figure)):
            next = v + 1
            if v == len(self.figure) - 1:
                next = 0
            if self.figure[v][1] > self.figure[next][1]:
                next = self.figure[next]
                v = self.figure[v]
            else:
                temp = v
                v = self.figure[next]
                next = self.figure[temp]

            rebra.append({'higher': v, 'lower': next})
        return rebra

    def fillFigure(self):
        self.clean_screen()
        self.paint_it()
        self.filter_it()

    def paint_it(self):
        rebra = sorted(self.getRebra(), key=lambda x: x.get('higher')[1], reverse=True)
        print(rebra)
        CAP = []
        for line in range(0, self.height):
            line = self.height - line

            CAP1 = []
            for r in CAP:
                if r.get('vertexes').get('lower')[1] < line:
                    CAP1.append(r)

            CAP = CAP1

            for y in rebra:
                if y.get('higher')[1] == line:
                    x1 = y.get('higher')[0]
                    y1 = y.get('higher')[1]
                    x2 = y.get('lower')[0]
                    y2 = y.get('lower')[1]
                    dx = (x1 - x2) / (y2+0.3 - y1)
                    CAP.append({'vertexes': y, 'dx': dx, 'xi': x1})
                    # print(CAP)

            cross = []
            for i in range(0, len(CAP)):
                xi = CAP[i].get('xi')
                dx = CAP[i].get('dx')
                cross.append(xi)
                CAP[i].update({'xi': xi + dx})
            cross.sort()
            if testing and line % 10 == 0:
                print(cross)
            for i in range(0, len(cross)):
                if i % 2 == 0:
                    continue
                for p in range(int(cross[i - 1]), int(cross[i])):
                    self.set_pixel(p, line)

    def set_pixel(self, x, y, brightness=255):
        brightness = min(255, max(0, int(brightness)))
        if 0 <= x < self.width and 0 <= y < self.height:
            index = y * self.width + x
            self.field[index] = brightness

    def filter_it(self):
        f = [self.field[i:i + self.width] for i in range(0, len(self.field), self.width)]
        f = [[x[0]] + x + [x[len(x) - 1]] for x in f]
        f = [f[0]] + f
        f = f + [f[len(f) - 1]]
        for i in range(0, self.height):
            for j in range(0, self.height):
                p = (f[i + 1][j] + f[i][j + 1] + f[i + 1][j + 2] + f[i + 2][j + 1] + f[i][j] + f[i + 2][j + 2] +
                     f[i][j + 2] + f[i + 2][j]) * 1 / 8 + f[i + 1][j + 1] * 1 / 3
                self.set_pixel(j, i, int(p))

    def clean_screen(self):
        self.field = [0] * (self.width * self.height)

    def run(self):
        while not glfw.window_should_close(self.window):
            self.display()
            glfw.poll_events()

        glfw.destroy_window(self.window)
        glfw.terminate()

    def display(self):
        glClear(GL_COLOR_BUFFER_BIT)

        if visible:
            glDrawPixels(self.width, self.height,
                         GL_BLUE, GL_UNSIGNED_BYTE,
                         self.field)

        glfw.swap_buffers(self.window)

    def mouseCallback(self, w, button, action, _):
        global visible

        if not visible and not testing:
            if action == glfw.PRESS and button == glfw.MOUSE_BUTTON_LEFT:
                x, y = glfw.get_cursor_pos(w)

                try:
                    len(self.figure)
                except Exception as e:
                    self.figure = []

                print([int(x), self.height - int(y)])
                self.figure.append([int(x), self.height - int(y)])

    def key_callback(self, _, key, scancode, action, mods):
        global visible

        if action == glfw.PRESS or action == glfw.REPEAT:
            if 48 < key < 58:
                d = int(key) - 48
                self.cutter["x_max"] = d * 50 + 450
                self.cutter["y_max"] = d * 50 + 450
                self.cutter["x_min"] = -d * 50 + 450
                self.cutter["y_min"] = -d * 50 + 450
            if key == glfw.KEY_ESCAPE:
                self.figure = []
                visible = False
            elif key == glfw.KEY_UP:
                f = [self.field[i:i + self.width] for i in range(0, len(self.field), self.width)]
                f = [[0] * 20 + x for x in f]
                for i in range(0, 20):
                    f.append([0] * (self.width + 10))
                self.field = list(chain.from_iterable(f))
                self.width += 20
                self.height += 20
                glfw.set_window_size(self.window, self.width, self.height)
            elif key == glfw.KEY_DOWN:
                self.width -= 20
                self.height -= 20
                glfw.set_window_size(self.window, self.width, self.height)
                self.fillFigure()
            elif key == glfw.KEY_ENTER:
                try:
                    len(self.figure)
                except Exception as e:
                    self.figure = []

                if len(self.figure) >= 3:
                    visible = True
                    # self.fillFigure()
                    self.cutFigure()
                    self.drawFigure(self.figure)

def main():
    Lab4(900, 900, "Lab 4").run()


if __name__ == "__main__":
    main()
