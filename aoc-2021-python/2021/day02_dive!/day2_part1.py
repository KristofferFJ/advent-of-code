import read_functions

position = [0, 0]

instructions = read_functions.read_as_string_pairs("input.txt")


def handleInstruction(instruction):
    direction = instruction[0]
    moves = int(instruction[1])
    if direction == "forward":
        position[0] += moves
    elif direction == "up":
        position[1] -= moves
    elif direction == "down":
        position[1] += moves


for instruction in instructions:
    handleInstruction(instruction)

print(position)
print(position[0]*position[1])