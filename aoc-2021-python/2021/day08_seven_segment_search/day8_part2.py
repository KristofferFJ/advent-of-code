import read_functions
import string_functions

input_output_numbers = [[j[0].split(), j[1].split()] for j in
                        [i.split("|") for i in read_functions.read_as_strings("input.txt")]]


class Wiring:
    def __init__(self, input_output):
        self.input = input_output[0]
        self.output = input_output[1]
        remaining_letters = self.input
        one = [i for i in self.input if len(i) == 2][0]
        remaining_letters.remove(one)
        seven = [i for i in self.input if len(i) == 3][0]
        remaining_letters.remove(seven)
        four = [i for i in self.input if len(i) == 4][0]
        remaining_letters.remove(four)
        eight = [i for i in self.input if len(i) == 7][0]
        remaining_letters.remove(eight)
        six = [i for i in self.input if len(i) == 6 and not string_functions.contains_letters(i, one)][0]
        remaining_letters.remove(six)
        nine = [i for i in self.input if len(i) == 6 and string_functions.contains_letters(i, four)][0]
        remaining_letters.remove(nine)
        zero = [i for i in remaining_letters if len(i) == 6][0]
        remaining_letters.remove(zero)
        five = [i for i in self.input if len(i) == 5 and string_functions.contains_letters(six, i)][0]
        remaining_letters.remove(five)
        three = [i for i in self.input if len(i) == 5 and string_functions.contains_letters(i, one)][0]
        remaining_letters.remove(three)
        two = remaining_letters[0]

        self.numbers = [zero, one, two, three, four, five, six, seven, eight, nine]

    def determine_output_numbers(self):
        wired_numbers = ""
        for number in self.output:
            wired_numbers += str(self.find_matching_number(number))
        return int(wired_numbers)

    def find_matching_number(self, number):
        for potential_match in self.numbers:
            if string_functions.has_same_letters(potential_match, number):
                return self.numbers.index(potential_match)


sum_of_output = 0
for input_output in input_output_numbers:
    wiring = Wiring(input_output)
    sum_of_output += wiring.determine_output_numbers()

print(sum_of_output)
