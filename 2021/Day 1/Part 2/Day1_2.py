import numpy as np

input = open("../input.txt").read().split()

increasements = 0
decreasements = 0

data = np.array(input)
data = data.astype(int)

for i in range(len(data)):
    try:
        if sum(data[i:i+3]) < sum(data[i+1:i+4]):
            increasements += 1
        elif sum(data[i:i+3]) > sum(data[i+1:i+4]):
            decreasements += 1
    except:
        break

print(f"increasements: {increasements}, decreasements: {decreasements}")