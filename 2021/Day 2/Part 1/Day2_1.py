input = open("../input.txt").read()

horizontal = 0
depth = 0

with open("../input.txt") as file:
    lines = file.readlines()
    lines = [line.rstrip() for line in lines]

for i in lines:
    element = str.split(i)
    if element[0] == "forward":
        horizontal += int(element[1])
    elif element[0] == "down":
        depth += int(element[1])
    elif element[0] == "up":
        depth -= int(element[1])

print(horizontal)
print(depth)

result = horizontal * depth

print(result)
