import read_functions

position = [0, 0, 0]
instructions = read_functions.read_as_string_pairs("input.txt")


def handle_instruction(_instruction):
    direction = _instruction[0]
    moves = int(_instruction[1])
    if direction == "forward":
        position[0] += moves
        position[1] += moves * position[2]
    elif direction == "up":
        position[2] -= moves
    elif direction == "down":
        position[2] += moves


for instruction in instructions:
    handle_instruction(instruction)

print(position)
print(position[0] * position[1])
