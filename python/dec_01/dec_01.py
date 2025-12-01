import sys
from typing import List

def first():
    with open("input_01.txt", "r") as input:
        lines = [line.strip() for line in input]

    result = 0
    dial = 50

    for line in lines:
        direction = line[0]
        number = int(line[1:])

        dial += number if direction == "R" else -number
        dial %= 100
        
        if dial == 0:
            result += 1

    print(f"1: {result}")

def second():
    with open("input_01.txt", "r") as input:
        lines = [line.strip() for line in input]

    result = 0
    dial = 50

    for line in lines:
        direction = line[0]
        number = int(line[1:])

        step = 1 if direction == "R" else -1
        
        for _ in range(number):
            dial = (dial + step) % 100
            if dial == 0:
                result += 1

    print(f"2: {result}")

first()
second()