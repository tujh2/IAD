
class FigureColor:
    def __init__(self):
        self._color = None
    
    @property
    def color(self):
        return self._color

    @color.setter
    def color(self, value):
        self._color = value

    @color.deleter
    def color(self):
        del self._color
