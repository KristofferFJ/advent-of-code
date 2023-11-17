import read_functions as read

list_of_ints = read.read_as_ints("input.txt")

sum = 0
line = 0
for i in range(3, len(list_of_ints)):
    if list_of_ints[i] + list_of_ints[i - 1] + list_of_ints[i - 2] > list_of_ints[i - 1] + list_of_ints[i - 2] + \
            list_of_ints[i - 3]:
        sum += 1

print(sum)
