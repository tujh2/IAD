from lab_python_oop.Rectangle import Rectangle

class Square(Rectangle):
    type = "Square"

    def __init__(self, s, c):
        super().__init__(s, s, c)
    
    def __repr__(self):
        return "{} side={}, area={}, color={}\n"\
                .format(
                    self.get_type(),
                    self.width, 
                    self.area(), 
                    self.color.color
                )
