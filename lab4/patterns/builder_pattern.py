#!/usr/bin/env python
from __future__ import annotations
from abc import ABC, abstractmethod, abstractproperty
from typing import Any


class ImageBuilderInterface(ABC):

    @abstractproperty
    def image(self) -> None:
        pass

    @abstractmethod
    def round(self) -> None:
        pass
    
    @abstractmethod
    def placeholder(self, uri) -> None:
        pass

    @abstractmethod
    def size(self, width, height) -> None:
        pass


class ImageBuilder(ImageBuilderInterface):

    def __init__(self) -> None:
        self.reset()

    def reset(self) -> None:
        self._image = Image()

    @property
    def image(self) -> Image:
        image = self._image
        self.reset()
        return image

    def round(self) -> None:
        self._image._is_round = True

    def placeholder(self, uri) -> None:
        self._image.placeholder = uri

    def size(self, width, height) -> None:
        self._image.set_size(width, height)


class Image():

    def __init__(self) -> None:
        self._width = 0
        self._height = 0
        self.placeholder = "empty"
        self._is_round = False

    def is_round(self):
        return self._is_round

    def set_size(self, width, height):
        self._width = width
        self._height = height

    def show_image(self):
        print("Image with placeholder {}, size {}x{}, isRound {}\n".format(self.placeholder, self._width, self._height, self._is_round))


class ImageDirector:

    def __init__(self) -> None:
        self._builder = None

    @property
    def builder(self) -> ImageBuilderInterface:
        return self._builder

    @builder.setter
    def builder(self, builder: ImageBuilderInterface) -> None:
        self._builder = builder

    def produce_standart_image(self) -> None:
        self.builder.size(220, 220)
        self.builder.placeholder("uri://placeholders/default_placeholder")

    def produce_rounded_image(self) -> None:
        self.builder.round()

    def produce_full_image(self, width, height, placeholder) -> None:
        self.builder.placeholder(placeholder)
        self.builder.size(width, height)


if __name__ == "__main__":

    director = ImageDirector()
    builder = ImageBuilder()
    director.builder = builder

    print("Standard default image: ")
    director.produce_standart_image()
    builder.image.show_image()

    print("Standard full featured product: ")
    director.produce_full_image(128, 128, "uri://placeholder/my_placeholder")
    builder.image.show_image()

    # can be used without director
    print("Custom image: ")
    builder.placeholder("uri://placeholder/custom_placeholder")
    builder.round()
    builder.size(300, 300)
    builder.image.show_image()
