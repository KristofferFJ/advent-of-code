import collection_functions
import read_functions


caves = set(list_functions.flatten_list([i.split("-") for i in read_functions.read_as_strings("input.txt")]))
paths = [i.split("-") for i in read_functions.read_as_strings("input.txt")]


def get_connections(cave):
    connections = []
    for path in paths:
        if path[0] == cave:
            connections.append(path[1])
        if path[1] == cave:
            connections.append(path[0])
    return connections


def find_all_paths():
    all_connections = []
    for connection in get_connections("start"):
        for path in continue_path(["start"], connection):
            all_connections.append(path)
    start_end_connections = [start_end_connection for start_end_connection in all_connections if start_end_connection[0] == "start" and start_end_connection[-1] == "end"]
    return start_end_connections


def continue_path(path, new_connection):
    new_path = [i for i in path]
    new_path.append(new_connection)
    if new_path[-1] == "end":
        return [new_path]
    if too_many_small_caves_visited(new_path):
        return [new_path]
    paths = []
    for connection in get_connections(new_path[-1]):
        paths.append(continue_path(new_path, connection))
    return list_functions.flatten_list(paths)


def too_many_small_caves_visited(path):
    small_caves_visited_twice = []
    for small_cave in [cave for cave in caves if not cave.isupper()]:
        visits_to_small_cave = len([cave_in_path for cave_in_path in path if cave_in_path == small_cave])
        if visits_to_small_cave > 2:
            return True
        if visits_to_small_cave > 1 and (small_cave == "start" or small_cave == "end"):
            return True
        if visits_to_small_cave > 1:
            small_caves_visited_twice.append(small_cave)
        if len(small_caves_visited_twice) > 1:
            return True
    return False


print(len(find_all_paths()))

