from PIL import Image
from colorthief import ColorThief
import numpy as np
import matplotlib.pyplot as plt

pinkArr2d = [[0 for i in range(21)] for j in range(21)]
redArr2d = [[0 for i in range(21)] for j in range(21)]
yellowArr2d = [[0 for i in range(21)] for j in range(21)]

total = 0


# pink (235, 0, 246)
def isPink(r, g, b):
    diff = abs(r - 235) + abs(g - 0) + abs(b - 246)
    return diff < 25


# red (233, 0, 2)
def isRed(r, g, b):
    diff = abs(r - 233) + abs(g - 0) + abs(b - 2)
    return diff < 25


# yellow (251, 255, 9)
def isYellow(r, g, b):
    diff = abs(r - 251) + abs(g - 255) + abs(b - 9)
    return diff < 25


for i in range(1, 1986):
    im = Image.open(fr"./png/out-{i}.png")
    px = im.load()

    ct = ColorThief(fr"./png/out-{i}.png")
    background = ct.get_color()

    pink = isPink(*background)
    red = isRed(*background)
    yellow = isYellow(*background)

    for x in range(21):
        for y in range(21):
            colour = px[y * 15 + 5, x * 15 + 5]
            if not isPink(*colour) and not isRed(*colour) and not isYellow(*colour):
                if pink:
                    pinkArr2d[y][x] += 1
                if red:
                    redArr2d[y][x] += 1
                if yellow:
                    yellowArr2d[y][x] += 1
                total += 1

    print(i)

print(np.matrix(pinkArr2d))
print(np.matrix(redArr2d))
print(np.matrix(yellowArr2d))
print(total)

while True:
    inp = input()
    if inp == "pink":
        plt.imshow(pinkArr2d, cmap='binary', interpolation='nearest')
        plt.show()

    if inp == "red":
        plt.imshow(redArr2d, cmap='binary', interpolation='nearest')
        plt.show()

    if inp == "yellow":
        plt.imshow(yellowArr2d, cmap='binary', interpolation='nearest')
        plt.show()
