def contains_letters(this, that):
    for letter in that:
        if letter not in this:
            return False
    return True


def has_same_letters(this, that):
    if len(this) != len(that):
        return False
    for letter in that:
        if letter not in this:
            return False
    return True
