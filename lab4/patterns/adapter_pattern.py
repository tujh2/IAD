#!/usr/bin/env python
import math


class Hole(object):
    def __init__(self, r):
        self.r = r

    def put(self, obj):
        try:
            if self.r >= obj.r:
                return True
            else:
                return False
        except AttributeError:
            print("Can not use this object!")


class Square(object):
    def __init__(self, x, h):
        self.x = x
        self.h = h


class SquareHoleAdapter(object):
    def __init__(self, sq_obj):
        self.sq_obj = sq_obj

    @property
    def r(self):
        return math.sqrt(2*(self.sq_obj.x**2))/2


if __name__ == "__main__":
    h1 = Hole(5)
    h2 = Hole(2)
    s1 = Square(5, 7)
    s2 = Square(3, 3)
    sa = SquareHoleAdapter(s2)

    print("Square(5,7) into Hole(5): ")
    print(h1.put(s1))
    print("\nSquare(5,7) into Hole(5) via adapter: ")
    print(h1.put(sa))
    print("\nSquare(3,3) into Hole(2): ")
    print(h2.put(sa))
