from lab_python_oop.Rectangle import Rectangle
from lab_python_oop.Circle import Circle
from lab_python_oop.Square import Square
import numpy

def main():
    rect = Rectangle(3, 8, "#0000FF")
    circle = Circle(5, "#00FF00")
    square = Square(10, "#FF0000")
    print("", rect, circle, square)
    print(numpy.linspace(0, 1, 5))

if __name__ == "__main__":
    main()
