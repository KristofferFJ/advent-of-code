import read_functions


class SyntaxLine:
    def __init__(self, line):
        self.line = line
        self.syntax_order = []
        self.value_parentheses = 1
        self.value_brackets = 2
        self.value_curly_brackets = 3
        self.value_than_signs = 4
        self.sum_of_missing_signs = 0

    def find_first_error(self):
        print("line: " + self.line)
        for sign in self.line:
            if not self.register_sign(sign):
                print("returns 0")
                return 0
        self.calculate_result()
        print("returns " + str(self.sum_of_missing_signs))
        return self.sum_of_missing_signs

    def calculate_result(self):
        while len(self.syntax_order) > 0:
            if self.syntax_order[-1] == "(":
                self.handle_sign(self.value_parentheses)
                continue
            if self.syntax_order[-1] == "[":
                self.handle_sign(self.value_brackets)
                continue
            if self.syntax_order[-1] == "{":
                self.handle_sign(self.value_curly_brackets)
                continue
            if self.syntax_order[-1] == "<":
                self.handle_sign(self.value_than_signs)

    def handle_sign(self, value):
        self.sum_of_missing_signs *= 5
        self.sum_of_missing_signs += value
        self.syntax_order.pop(-1)

    def register_sign(self, sign):
        if sign in ("(", "[", "{", "<"):
            self.syntax_order.append(sign)
            return 1
        else:
            if sign == ")":
                if self.syntax_order[-1] == "(":
                    self.syntax_order.pop(-1)
                    return 1
            if sign == "]":
                if self.syntax_order[-1] == "[":
                    self.syntax_order.pop(-1)
                    return 1
            if sign == "}":
                if self.syntax_order[-1] == "{":
                    self.syntax_order.pop(-1)
                    return 1
            if sign == ">":
                if self.syntax_order[-1] == "<":
                    self.syntax_order.pop(-1)
                    return 1
        return 0


class SyntaxPage:
    def __init__(self, syntax_lines):
        self.syntax_lines = syntax_lines

    def calculate_syntax_error(self):
        sum_of_errors = []
        for line in self.syntax_lines:
            sum_of_errors.append(line.find_first_error())
        sum_of_errors = [value for value in sum_of_errors if value != 0]
        sum_of_errors.sort()

        return sum_of_errors[int(len(sum_of_errors) / 2)]


_syntax_lines = [SyntaxLine(line) for line in read_functions.read_as_strings("input.txt")]
syntaxPage = SyntaxPage(_syntax_lines)

print(syntaxPage.calculate_syntax_error())
