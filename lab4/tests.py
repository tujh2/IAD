#!/usr/bin/env python
from patterns.builder_pattern import ImageDirector, ImageBuilder
from patterns.adapter_pattern import Hole, Square, SquareHoleAdapter
from mock import patch

if __name__ == "__main__":
    director = ImageDirector()
    builder = ImageBuilder()
    director.builder = builder

    h1 = Hole(5)
    h2 = Hole(2)
    s1 = Square(5, 7)
    s2 = Square(3, 3)
    sa = SquareHoleAdapter(s2)

    def tdd():
        director.produce_standart_image()
        assert builder.image.placeholder == "uri://placeholders/default_placeholder"

        builder.round()
        assert builder.image.is_round() == True
        
        assert type(h1.put(sa)) == type(True)

    
    def mock():
        with patch("patterns.builder_pattern.Image.is_round") as m:
            builder.round()
            m.return_value = builder.image.is_round()

        with patch("patterns.adapter_pattern.Hole.put") as m:
            m.return_value = h1.put(sa)

    tdd()
    mock()

