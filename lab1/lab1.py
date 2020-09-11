#!/usr/bin/python
import sys
import math

class Coefficient:
    def __init__(self, letter):
        self.letter = letter
    
    value = float('nan')

    def __str__(self):
        return self.letter + " = " + str(self.value)


def readValue(coefficient):
    while(math.isnan(coefficient.value)):
        print("Enter coefficient " + coefficient.letter + ":")
        try:
            coefficient.value = float(input())
        except ValueError:
            print("Error! You've typed invalid coefficient " + coefficient.letter + "! Try again!")

def parseValue(coefficient, value):
    try:
        coefficient.value = float(value)
    except ValueError:
        print("Error! You've typed invalid coefficient " + coefficient.letter + "! Try again!")
        exit(0)


print("Разработал Сысойкин Егор; ИУ5-54Б")


a = Coefficient("A")
b = Coefficient("B")
c = Coefficient("C")

args = sys.argv
argsLen = len(args)

if argsLen == 1 or argsLen == 0:
    readValue(a)
    readValue(b)
    readValue(c)
elif argsLen == 2:
    parseValue(a, args[1])
    readValue(b)
    readValue(c)
elif argsLen == 3:
    parseValue(a, args[1])
    parseValue(b, args[2])
    readValue(c)
else:
    parseValue(a, args[1])
    parseValue(b, args[2])
    parseValue(c, args[3])
    
print(a, b, c)
print()

d = b.value*b.value - 4*a.value*c.value

if d < 0:
    print("There is no any real roots! (D < 0)")
elif a.value == 0:
    if b.value == 0:
        if c.value == 0:
            print("Anything")
        else:
            print("No roots")
    else:
        if -c.value / b.value >= 0:
            x = math.sqrt(-c.value/b.value)
            print("Two roots: x1 = ", x, "; x2 = ", -x)
        else:
            print("There is no any real roots!")
else:
    sqd = math.sqrt(d)
    y1 = (-b.value + sqd)/(2*a.value)
    y2 = (-b.value - sqd)/(2*a.value)
    
    if y1 < 0 and y2 < 0:
        print("There is no any real roots!")
        exit(0)

    if y1 >= 0:
        x = math.sqrt(y1)
        print ("x1 = ", x, "; x2 = ", -x)
    if y2 >= 0 and d != 0:
        x = math.sqrt(y2)
        print ("x3 = ", x, "; x4 = ", -x)

