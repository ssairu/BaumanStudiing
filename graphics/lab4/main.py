from itertools import chain

import glfw
from OpenGL.GL import *

visible = True
testing = True


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

        if testing:
            self.figure = [
                [10, 20], [250, 200], [300, 60], [400, 300], [300, 400]
            ]

        glfw.make_context_current(self.window)
        glfw.set_key_callback(self.window, self.key_callback)
        glfw.set_mouse_button_callback(self.window, self.mouseCallback)

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
                    dx = (x1 - x2) / (y2 - y1)
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

                self.figure.append([int(x), self.height - int(y)])

    def key_callback(self, _, key, scancode, action, mods):
        global visible

        if action == glfw.PRESS or action == glfw.REPEAT:
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
                    self.fillFigure()


def main():
    Lab4(800, 800, "Lab 4").run()


if __name__ == "__main__":
    main()
