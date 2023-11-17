def flatten_list(nested_lists):
    flat_list = []
    for sub_list in nested_lists:
        for element in sub_list:
            flat_list.append(element)
    return flat_list


def intersect_ranges(x, y):
    return range(max(x[0], y[0]), min(x[-1], y[-1]) + 1)