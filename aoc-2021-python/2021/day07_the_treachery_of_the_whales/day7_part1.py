import read_functions

list_of_crabs_starting_position = [int(i) for i in read_functions.read_as_strings("input.txt")[0].split(",")]

highest_starting_position = max(list_of_crabs_starting_position)
lowest_starting_position = min(list_of_crabs_starting_position)

current_best_value = lowest_starting_position

def calculation_fuel_usage(position, crab_positions):
    fuel_usage = 0
    for crab_position in crab_positions:
        fuel_usage += abs(crab_position - position)
    return fuel_usage

least_fuel_usage = calculation_fuel_usage(lowest_starting_position, list_of_crabs_starting_position)
position = lowest_starting_position

for common_position in range(lowest_starting_position, highest_starting_position):
    fuel_usage = calculation_fuel_usage(common_position, list_of_crabs_starting_position)
    if fuel_usage < least_fuel_usage:
        least_fuel_usage = fuel_usage
        position = common_position

print(position)
print(least_fuel_usage)

