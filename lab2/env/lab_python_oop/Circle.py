from lab_python_oop.Figure import Figure
from lab_python_oop.FigureColor import FigureColor
from math import pi

class Circle(Figure):
    type = "Circle"

    @classmethod
    def get_type(cls):
        return cls.type

    def __init__(self, r, c):
        self.radius = r
        self.color = FigureColor()
        self.color.color = c

    def area(self):
        return 2 * pi * self.radius * self.radius

    def __repr__(self):
        return "{} radius={}, area={}, color={}\n"\
                .format(
                    self.get_type(),
                    self.radius,
                    self.area(),
                    self.color.color
                )
