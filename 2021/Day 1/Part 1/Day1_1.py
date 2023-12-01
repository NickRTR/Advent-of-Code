import numpy as np

input = open("../input.txt").read().split()

increasements = 0
decreasements = 0

data = np.array(input)
data = data.astype(int)

result = np.diff(data)

for i in result:
    if i > 0:
        increasements += 1
    elif i < 0:
        decreasements += 1

print(f"increasements: {increasements}, decreasements: {decreasements}")