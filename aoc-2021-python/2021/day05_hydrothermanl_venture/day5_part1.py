import read_functions as read
import TableIt

# [[x1, y1], [x2, y2]] x and y as ints.
start_end_points = [[i[0].split(","), i[1].split(",")] for i in
                    [i.split(" -> ") for i in read.read_as_strings("input.txt")]]


covered_points = [[0 for i in range(1000)] for j in range(1000)]

# takes points [[x1, y1], [x2, y2]]
def add_covered_points(point):
    x1 = int(point[0][0])
    y1 = int(point[0][1])
    x2 = int(point[1][0])
    y2 = int(point[1][1])

    if x1 == x2:
        largest_value = max(y1, y2)
        smallest_value = min(y1, y2)
        for point in [[x1, i] for i in range(smallest_value, largest_value + 1)]:
            covered_points[point[0]][999 - point[1]] += 1
    if y1 == y2:
        largest_value = max(x1, x2)
        smallest_value = min(x1, x2)
        for point in [[i, y1] for i in range(smallest_value, largest_value + 1)]:
            covered_points[point[0]][999 - point[1]] += 1
        return
    return []


for point in start_end_points:
    add_covered_points(point)

twice_hit_count = 0
for row in covered_points:
    for element in row:
        if element >= 2:
            twice_hit_count += 1

TableIt.printTable(covered_points)
print(twice_hit_count)