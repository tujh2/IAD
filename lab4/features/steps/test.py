from behave import *
from patterns.builder_pattern import ImageDirector, ImageBuilder

@given("build standart image")
def step_impl(context):
    context.director = ImageDirector()
    context.builder = ImageBuilder()
    context.director.builder = context.builder

@then("image is standart")
def step_impl(context):
    context.director.produce_standart_image()
    image = context.builder.image
    assert image.placeholder == "uri://placeholders/default_placeholder"
    assert image.is_round() == False
