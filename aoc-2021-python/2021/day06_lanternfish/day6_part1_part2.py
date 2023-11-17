# [0, 1, 2, 3, 4, 5, 6, 7, 8]
import read_functions

starter_fish = read_functions.read_as_strings("input.txt")[0].split(",")
lantern_fish = [0 for i in range(9)]
for i in range(9):
    specific_starter_fish = [int(j) for j in starter_fish if int(j) == i]
    lantern_fish[i] = len(specific_starter_fish)

print(starter_fish)
print(lantern_fish)

def pass_day():
    zero_day = lantern_fish[0]
    one_day = lantern_fish[1]
    two_day = lantern_fish[2]
    three_day = lantern_fish[3]
    four_day = lantern_fish[4]
    five_day = lantern_fish[5]
    six_day = lantern_fish[6]
    seven_day = lantern_fish[7]
    eight_day = lantern_fish[8]

    lantern_fish[0] = one_day
    lantern_fish[1] = two_day
    lantern_fish[2] = three_day
    lantern_fish[3] = four_day
    lantern_fish[4] = five_day
    lantern_fish[5] = six_day
    lantern_fish[6] = seven_day + zero_day
    lantern_fish[7] = eight_day
    lantern_fish[8] = zero_day

for i in range(80):
    pass_day()

print(sum(lantern_fish))

for i in range(256 - 80):
    pass_day()

print(sum(lantern_fish))