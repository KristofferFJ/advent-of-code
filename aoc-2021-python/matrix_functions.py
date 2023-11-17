def transpose_matrix(m):
    return [[m[j][i] for j in range(len(m))] for i in range(len(m[0]))]


def prod(array):
    product = 1
    for i in array:
        product *= i
    return product