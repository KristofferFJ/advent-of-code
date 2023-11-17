import read_functions
import collection_functions


class Instruction:
    def __init__(self, _line):
        _split_line = _line.split(" ")
        _coordinate_instructions = _split_line[1].split(",")
        x_instructions = _coordinate_instructions[0][2:].split("..")
        y_instructions = _coordinate_instructions[1][2:].split("..")
        z_instructions = _coordinate_instructions[2][2:].split("..")
        self.x_range = range(int(x_instructions[0]), int(x_instructions[1]) + 1)
        self.y_range = range(int(y_instructions[0]), int(y_instructions[1]) + 1)
        self.z_range = range(int(z_instructions[0]), int(z_instructions[1]) + 1)
        self.toggle = True if _split_line[0] == "on" else False

    def execute(self, _grid):
        for x in self.x_range:
            for y in self.y_range:
                for z in self.z_range:
                    _grid.update(x, y, z, self.toggle)

    def execute_with_boundaries(self, _grid, boundary_range):
        for x in collection_functions.intersect_ranges(self.x_range, boundary_range):
            for y in collection_functions.intersect_ranges(self.y_range, boundary_range):
                for z in collection_functions.intersect_ranges(self.z_range, boundary_range):
                    _grid.update(x, y, z, self.toggle)


class Grid3:
    def __init__(self, _range, initial_position):
        self.range = _range
        self.grid = [[[initial_position for i in self.range] for j in self.range] for k in self.range]

    def update(self, x, y, z, new_value):
        self.grid[x][y][z] = new_value

    def get(self, x, y, z):
        return self.grid[x][y][z]

    def count_points_with_value(self, value):
        return self.count_points_with_value_in_range(value, self.range)

    def count_points_with_value_in_range(self, value, _range):
        points_with_value = 0
        for x in _range:
            for y in _range:
                for z in _range:
                    if self.get(x, y, z) == value:
                        points_with_value += 1
        return points_with_value


input_lines = read_functions.read_as_strings("input.txt")
instructions = [Instruction(input_line) for input_line in input_lines]
active_range = range(-50, 50 + 1)
grid = Grid3(active_range, False)

for _instruction in instructions:
    _instruction.execute_with_boundaries(grid, active_range)

print(grid.count_points_with_value(True))
