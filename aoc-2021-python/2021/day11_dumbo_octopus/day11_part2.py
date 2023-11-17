import read_functions


class Octopus:
    def __init__(self, energy_level, row, column):
        self.energy_level = int(energy_level)
        self.flashed = False
        self.row = row
        self.column = column

    def up_energy_level(self):
        self.energy_level += 1

    def prepare_for_next_step(self):
        if self.energy_level > 9:
            self.energy_level = 0
        self.flashed = False

    def flash(self):
        if self.energy_level > 9 and not self.flashed:
            # print("Octopus at i=" + str(self.row) + ", j=" + str(self.column) + " flashed")
            self.flashed = True
            return 1
        return 0


class OctopusGrid:
    def __init__(self, _octopi):
        self.grid = _octopi
        self.total_flashes = 0

    def take_step(self):
        self.up_energy_levels()
        flashed_in_step = self.flash_step()
        self.prepare_for_next_step()
        return flashed_in_step

    def flash_step(self):
        flashed_in_step = 0
        flashed_in_last_round = self.flash_all()
        while flashed_in_last_round:
            flashed_in_step += flashed_in_last_round
            flashed_in_last_round = self.flash_all()
        return flashed_in_step

    def flash_all(self):
        flashed_octopi = 0
        for line in self.grid:
            for octopus in line:
                if octopus.flash():
                    flashed_octopi += 1
                    for neighbor_octopus in self.get_neighbors(octopus):
                        neighbor_octopus.up_energy_level()
        self.total_flashes += flashed_octopi
        return flashed_octopi

    def get_neighbors(self, octopus):
        neighbors = []
        i = octopus.row
        j = octopus.column
        max_row = len(self.grid) - 1
        max_column = len(self.grid[0]) - 1
        # 1 8 7
        # 2   6
        # 3 4 5
        if i > 0 and j > 0:  # 1
            neighbors.append(self.grid[i - 1][j - 1])
        if j > 0:  # 2
            neighbors.append(self.grid[i][j - 1])
        if j > 0 and i < max_row:  # 3
            neighbors.append(self.grid[i + 1][j - 1])
        if i < max_row:  # 4
            neighbors.append(self.grid[i + 1][j])
        if i < max_row and j < max_column:  # 5
            neighbors.append(self.grid[i + 1][j + 1])
        if j < max_column:  # 6
            neighbors.append(self.grid[i][j + 1])
        if j < max_column and i > 0:  # 7
            neighbors.append(self.grid[i - 1][j + 1])
        if i > 0:  # 8
            neighbors.append(self.grid[i - 1][j])
        return neighbors

    def up_energy_levels(self):
        for line in self.grid:
            for octopus in line:
                octopus.up_energy_level()

    def prepare_for_next_step(self):
        for line in self.grid:
            for octopus in line:
                octopus.prepare_for_next_step()


input_rows = read_functions.read_as_strings("input.txt")
_octopi = [[Octopus(input_rows[i][j], i, j) for j in range(len(input_rows[i]))] for i in range(len(input_rows))]
_octopusGrid = OctopusGrid(_octopi)

i = 0
while True:
    i += 1
    flashed_in_step = _octopusGrid.take_step()
    print("Flashed in Step " + str(i) + ": " + str(flashed_in_step))
    if flashed_in_step == 100:
        break
