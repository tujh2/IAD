from lab_python_oop.Figure import Figure
from lab_python_oop.FigureColor import FigureColor

class Rectangle(Figure):
    type="Rectangle"

    @classmethod
    def get_type(cls):
        return cls.type
    
    def __init__(self, w, h, c):
        self.width = w
        self.height = h
        self.color = FigureColor()
        self.color.color = c

    def area(self):
        return self.width * self.height

    def __repr__(self):
        return "{} width={}, height={}, area={}, color={}\n"\
                .format(
                    self.get_type(),
                    self.width,
                    self.height,
                    self.area(),
                    self.color.color
                )

