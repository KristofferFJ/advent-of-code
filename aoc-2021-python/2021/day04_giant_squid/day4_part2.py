import read_functions as read
import matrix_functions as m

list_of_strings = read.read_as_strings("input.txt")
drawn_numbers = [int(i) for i in list_of_strings[0].split(",")]


def create_boards():
    boards = []
    current_board = []
    for row in range(2, len(list_of_strings)):
        if list_of_strings[row] == "":
            boards.append([False, current_board])
            current_board = []
        else:
            current_board.append([[int(i), False] for i in list_of_strings[row].split()])
    boards.append([False, current_board])
    return boards


def mark_board(board, number):
    for row in board:
        for element in row:
            if element[0] == number:
                element[1] = True


def mark_boards(number):
    for board in boards:
        mark_board(board[1], number)


def check_for_success(board):
    for row in board:
        if is_row_marked(row):
            return True
    for column in m.transpose_matrix(board):
        if is_row_marked(column):
            return True
    return False


def mark_winners():
    for board in boards:
        if check_for_success(board[1]):
            board[0] = True


def is_row_marked(row):
    for element in row:
        if not element[1]:
            return False
    return True


def check_for_last_loser():
    losers = 0
    loser_board = []
    for board in boards:
        if not board[0]:
            losers += 1
            loser_board = board[1]
    if losers == 1:
        return loser_board


def all_boards_done():
    for board in boards:
        if not board[0]:
            return False
    return True


boards = create_boards()
loser_board = []
for draw_number in drawn_numbers:
    mark_boards(draw_number)
    mark_winners()
    if check_for_last_loser():
        loser_board = check_for_last_loser()
    if all_boards_done():
        sum_of_false_elements = 0
        for row in loser_board:
            for element in row:
                if not element[1]:
                    sum_of_false_elements += element[0]
        print(sum_of_false_elements)
        print(sum_of_false_elements * draw_number)
        break