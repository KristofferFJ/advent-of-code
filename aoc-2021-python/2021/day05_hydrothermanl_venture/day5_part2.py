import read_functions as read
import TableIt

# [[x1, y1], [x2, y2]] x and y as ints.
start_end_points = [[i[0].split(","), i[1].split(",")] for i in
                    [i.split(" -> ") for i in read.read_as_strings("input.txt")]]

covered_points = [[0 for i in range(1000)] for j in range(1000)]


# takes points [[x1, y1], [x2, y2]]
def add_covered_points(start_end_point):
    for covered_point in get_covered_point(start_end_point):
        covered_points[covered_point[0]][covered_point[1]] += 1


def is_slope_upwards(point):
    x_movement = int(point[0][0]) - int(point[1][0])
    y_movement = int(point[0][1]) - int(point[1][1])
    return x_movement * y_movement > 0


def get_covered_point(start_end_point):
    x1 = int(start_end_point[0][0])
    y1 = int(start_end_point[0][1])
    x2 = int(start_end_point[1][0])
    y2 = int(start_end_point[1][1])

    max_x = max(x1, x2)
    min_x = min(x1, x2)
    max_y = max(y1, y2)
    min_y = min(y1, y2)

    if x1 == x2:
        return [[x1, i] for i in range(min_y, max_y + 1)]
    if y1 == y2:
        return [[i, y1] for i in range(min_x, max_x + 1)]
    if max_x - min_x == max_y - min_y:
        if is_slope_upwards(start_end_point):
            return [[min_x + i, min_y + i] for i in range(max_y - min_y + 1)]
        else:
            return [[min_x + i, max_y - i] for i in range(max_y - min_y + 1)]
    return []


for point in start_end_points:
    add_covered_points(point)

twice_hit_count = 0
for row in covered_points:
    for element in row:
        if element >= 2:
            twice_hit_count += 1

print(twice_hit_count)
