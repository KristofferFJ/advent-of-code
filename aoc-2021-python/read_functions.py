def read_as_ints(filename):
    f = open(filename, "r")
    string_list = f.read().split("\n")
    return [int(i) for i in string_list]


def read_as_strings(filename):
    return open(filename, "r").read().split("\n")


def read_as_string_pairs(filename):
    f = open(filename, "r")
    string_list = f.read().split("\n")
    return [i.split(" ") for i in string_list]
