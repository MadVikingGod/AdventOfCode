gameInput = """325489"""
from enum import Enum

def buildSpiral(num):
    max = 0
    size = 1
    x = 1
    while max < num:
        x+=1
        max += size *4
        size += 2
        print(max)
        print(size)
        print(x)
    x -=1


buildSpiral(int(gameInput))
class Loc(object):
    def __init__(self, x=0, y=0, step=0, steps=0, maxSteps=0):
        self.x = x
        self.y = y
        self.step = step
        self.steps = steps
        self.maxSteps = maxSteps

    def next(self):
        steps = self.steps +1
        if self.step == 0:
            if steps<self.maxSteps+1:
                return Loc(x=self.x, y=self.y+1, step=self.step, steps=steps, maxSteps=self.maxSteps)
            return Loc(x=self.x, y=self.y+1, step=1, maxSteps=self.maxSteps+1)
        elif self.step == 1:
            if steps<self.maxSteps:
                return Loc(x=self.x-1, y=self.y, step=self.step, steps=steps, maxSteps=self.maxSteps)
            return Loc(x=self.x-1, y=self.y, step=2, maxSteps=self.maxSteps)
        elif self.step == 2:
            if steps<self.maxSteps+1:
                return Loc(x=self.x, y=self.y-1, step=self.step, steps=steps, maxSteps=self.maxSteps)
            return Loc(x=self.x, y=self.y-1, step=3, maxSteps=self.maxSteps+1)
        elif self.step == 3:
            if steps<self.maxSteps:
                return Loc(x=self.x+1, y=self.y, step=self.step, steps=steps, maxSteps=self.maxSteps)
            return Loc(x=self.x+1, y=self.y, step=0, maxSteps=self.maxSteps)



def adjecents(loc, grid):
    # print(loc.x, loc.y, loc.step, loc.steps, loc.maxSteps)
    xs = [loc.x+1, loc.x, loc.x-1]
    ys = [loc.y+1, loc.y, loc.y-1]
    t = [grid[x][y] for x in xs for y in ys ]
    return sum(t)

X = 11
Y = 11
grid = []
grid = [[0]*X for y in range(Y)]
loc = Loc(x=X//2, y=Y//2)
grid[loc.x][loc.y] = 1

while loc.x>=0 and loc.y>=0:
    loc = loc.next()
    if loc.x <0 or loc.y < 0 or loc.x>=X-1 or loc.y>=Y-1:
        for row in grid:
            print(row)
    ad = adjecents(loc,grid)
    if ad > int(gameInput):
        print("won",ad)
    grid[loc.x][loc.y] = ad

for row in grid:
    print(row)


