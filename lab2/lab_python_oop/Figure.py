from abc import ABCMeta,abstractmethod

class Figure(metaclass=ABCMeta):

    @abstractmethod
    def area(self):
        pass
