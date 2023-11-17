import read_functions
import timeit

start = timeit.default_timer()


heights = [[int(j) for j in i] for i in read_functions.read_as_strings("input.txt")]
columns = len(heights[0]) - 1
rows = len(heights) - 1

low_points = []


def is_low_point(i, j):
    height_ij = heights[i][j]
    if i == 0 and j == 0:
        if height_ij < heights[i + 1][j] and height_ij < heights[i][j + 1]:
            return True
        return False
    if i == 0 and j == columns:
        if height_ij < heights[i + 1][j] and height_ij < heights[i][j - 1]:
            return True
        return False
    if i == heights and j == 0:
        if height_ij < heights[i][j + 1] and height_ij < heights[i - 1][j]:
            return True
        return False
    if i == heights and j == columns:
        if height_ij < heights[i - 1][j] and height_ij < heights[i][j - 1]:
            return True
        return False
    if i == 0:
        if height_ij < heights[i][j - 1] and height_ij < heights[i][j + 1] and height_ij < heights[i + 1][j]:
            return True
        return False
    if j == 0:
        if height_ij < heights[i - 1][j] and height_ij < heights[i + 1][j] and height_ij < heights[i][j + 1]:
            return True
        return False
    if i == rows:
        if height_ij < heights[i][j - 1] and height_ij < heights[i][j + 1] and height_ij < heights[i - 1][j]:
            return True
        return False
    if j == columns:
        if height_ij < heights[i - 1][j] and height_ij < heights[i + 1][j] and height_ij < heights[i][j - 1]:
            return True
        return False
    if height_ij < heights[i - 1][j] and height_ij < heights[i + 1][j] and height_ij < heights[i][j - 1] and height_ij < \
            heights[i][j + 1]:
        return True
    return False


for i in range(rows + 1):
    for j in range(columns + 1):
        if is_low_point(i, j):
            low_points.append(heights[i][j])

print(low_points)
print(sum([i + 1 for i in low_points]))

print(timeit.default_timer() - start)