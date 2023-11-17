import read_functions as read

list_of_ints = read.read_as_strings("input.txt")

zeros = [0 for i in list_of_ints[0]]
ones = [0 for i in list_of_ints[0]]


def find_most_common_digit(index):
    zeros = 0
    ones = 0
    for line in list_of_ints:
        if line[index] == "0":
            zeros += 1
        else:
            ones += 1
    if zeros > ones:
        return 0
    return 1


def bit_array_to_dec(array):
    exp = 0
    value = 0
    for i in range(1, len(array) + 1):
        value += array[-i] * 2 ** exp
        exp += 1
    return value


gamma = [find_most_common_digit(i) for i in range(len(list_of_ints[0]))]
delta = [(i + 1) % 2 for i in gamma]

print(gamma)
print(delta)

print(bit_array_to_dec(gamma))
print(bit_array_to_dec(delta))

print(bit_array_to_dec(gamma) * bit_array_to_dec(delta))
