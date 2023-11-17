import read_functions

one_length = 2
four_length = 4
seven_length = 3
eight_length = 7

input = read_functions.read_as_strings("input.txt")
input_output_numbers = [[j[0].split(), j[1].split()] for j in [i.split("|") for i in input]]

sum_recognized_output = 0
for line in input_output_numbers:
    for number_string in line[1]:
        if len(number_string) in (one_length, four_length, seven_length, eight_length):
            sum_recognized_output += 1

print(sum_recognized_output)
