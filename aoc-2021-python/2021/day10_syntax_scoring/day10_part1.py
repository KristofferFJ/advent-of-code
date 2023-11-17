import read_functions


class SyntaxLine:
    def __init__(self, line):
        self.line = line
        self.syntax_order = []
        self.value_parentheses = 3
        self.value_brackets = 57
        self.value_curly_brackets = 1197
        self.value_than_signs = 25137

    def find_first_error(self):
        print("line: " + self.line)
        for sign in self.line:
            if not self.register_sign(sign):
                print("returns " + str(self.calculate_result(sign)))
                return self.calculate_result(sign)
        print("returns 0")
        return 0

    def calculate_result(self, sign):
        if sign == ")":
            return self.value_parentheses
        if sign == "]":
            return self.value_brackets
        if sign == "}":
            return self.value_curly_brackets
        if sign == ">":
            return self.value_than_signs

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
        sum_of_errors = 0
        for line in self.syntax_lines:
            sum_of_errors += line.find_first_error()
        return sum_of_errors


_syntax_lines = [SyntaxLine(line) for line in read_functions.read_as_strings("input.txt")]
syntaxPage = SyntaxPage(_syntax_lines)

print(syntaxPage.calculate_syntax_error())