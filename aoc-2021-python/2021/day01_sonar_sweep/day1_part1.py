import read_functions as read

list_of_ints = read.read_as_ints("input.txt")

sum = 0
line = 0
for i in range(1, len(list_of_ints)):
    line += 1
    if list_of_ints[i] > list_of_ints[i-1]:
        sum += 1

print(sum)
