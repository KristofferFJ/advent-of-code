import matrix_functions
import read_functions

heights = [[int(j) for j in i] for i in read_functions.read_as_strings("input.txt")]
last_column = len(heights[0]) - 1
last_row = len(heights) - 1


def get_adjacent_and_itself(i, j):
    adjacent = [[i, j]]
    if i != 0:
        adjacent.append([i - 1, j])
    if j != 0:
        adjacent.append([i, j - 1])
    if i != last_row:
        adjacent.append([i + 1, j])
    if j != last_column:
        adjacent.append([i, j + 1])
    return adjacent


class Pool:
    def __init__(self, _heights):
        self.heights = _heights
        self.pools = [[0 for _ in range(last_column + 1)] for _ in range(last_row + 1)]
        self.next_pool_number = 10

    def get_pool_number(self, i, j):
        for array in get_adjacent_and_itself(i, j):
            if self.pools[array[0]][array[1]] not in (0, 9):
                return self.pools[array[0]][array[1]]
        self.next_pool_number += 1
        return self.next_pool_number

    def mark_adjacent_and_itself(self, i, j, pool_number):
        if self.heights[i][j] == 9:
            return
        for array in get_adjacent_and_itself(i, j):
            pool_height = self.heights[array[0]][array[1]]
            if pool_height != 9:
                self.pools[array[0]][array[1]] = pool_number

    def assign_pools(self):
        for i in range(len(self.pools)):
            for j in range(len(self.pools[0])):
                pool_number = self.get_pool_number(i, j)
                self.mark_adjacent_and_itself(i, j, pool_number)

    def find_pools(self):
        return [self.count_basin_size(i) for i in range(1, self.next_pool_number + 1)]

    def count_basin_size(self, basin_number):
        items_in_basin = 0
        for row in self.pools:
            for element in row:
                if element == basin_number:
                    items_in_basin += 1
        return items_in_basin

    def reassign_numbers(self):
        changed_values = 0
        for i in range(len(self.pools)):
            for j in range(len(self.pools[0])):
                pool_value = self.pools[i][j]
                if pool_value == 0:
                    continue
                lowest_neighbor = min([self.pools[array[0]][array[1]] for array in get_adjacent_and_itself(i, j) if self.pools[array[0]][array[1]] != 0])
                if pool_value != lowest_neighbor:
                    print("changes " + str(pool_value) + " to " + str(lowest_neighbor))
                    self.pools[i][j] = lowest_neighbor
                    changed_values += 1
        if changed_values != 0:
            self.reassign_numbers()


pool = Pool(heights)
pool.assign_pools()
pool.reassign_numbers()

print(pool.pools)
print(sorted(pool.find_pools(), reverse=True)[:3])
print(matrix_functions.prod(sorted(pool.find_pools(), reverse=True)[:3]))
