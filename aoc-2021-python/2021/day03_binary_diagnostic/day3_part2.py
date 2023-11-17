import read_functions as read

list_of_ints = read.read_as_strings("input.txt")
bit_number_length = len(list_of_ints[0])


def find_most_common_digit(bits, index, tiebreaker):
    zeros = 0
    ones = 0
    for line in bits:
        if line[index] == "0":
            zeros += 1
        else:
            ones += 1
    if zeros == ones:
        return tiebreaker
    if zeros > ones:
        return 0
    return 1


def find_least_common_digit(bits, index, tiebreaker):
    zeros = 0
    ones = 0
    for line in bits:
        if line[index] == "0":
            zeros += 1
        else:
            ones += 1
    if zeros == ones:
        return tiebreaker
    if zeros > ones:
        return 1
    return 0


def bit_array_to_dec(array):
    exp = 0
    value = 0
    for i in range(1, len(array) + 1):
        value += int(array[-i]) * 2 ** exp
        exp += 1
    return value


def get_oxygen():
    remaining_bits = list_of_ints
    for index in range(bit_number_length):
        most_common_digit = find_most_common_digit(remaining_bits, index, 1)
        remaining_bits = [i for i in remaining_bits if i[index] == str(most_common_digit)]
        if len(remaining_bits) == 1:
            return remaining_bits[0]
    return remaining_bits[0]


def get_scrubber():
    remaining_bits = list_of_ints
    for index in range(bit_number_length):
        least_common_digit = find_least_common_digit(remaining_bits, index, 0)
        remaining_bits = [i for i in remaining_bits if i[index] == str(least_common_digit)]
        if len(remaining_bits) == 1:
            return remaining_bits[0]
    return remaining_bits[0]


print(get_oxygen())
print(get_scrubber())

print(bit_array_to_dec(get_oxygen()) * bit_array_to_dec(get_scrubber()))


