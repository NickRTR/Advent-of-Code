input = open("../input.txt").read()

horizontal = 0
depth = 0
aim = 0

with open("../input.txt") as file:
    lines = file.readlines()
    lines = [line.rstrip() for line in lines]

for i in lines:
    element = str.split(i)
    if element[0] == "forward":
        horizontal += int(element[1])
        depth += aim * int(element[1])
    elif element[0] == "down":
        aim += int(element[1])
    elif element[0] == "up":
        aim -= int(element[1])

print(horizontal)
print(depth)
print(aim)

result = horizontal * depth

print(result)
