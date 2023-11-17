class Player:
    def __init__(self, starting_position):
        self.position = starting_position
        self.score = 0

    def roll(self, _current_roll):
        print("current roll: " + str(current_roll))
        self.position = (self.position + (_current_roll * 3 + 2)) % 10 + 1
        print("now at: " + str(self.position))
        _current_roll += 3
        self.score += self.position
        print("score of: " + str(self.score))
        return _current_roll

    def has_won(self):
        if self.score > 999:
            return True
        return False


player1 = Player(3)
player2 = Player(10)

current_roll = 1
losing_player = player1
while True:
    current_roll = player1.roll(current_roll)
    if player1.has_won():
        losing_player = player2
        break
    current_roll = player2.roll(current_roll)
    if player2.has_won():
        losing_player = player1
        break

print("last number rolled: " + str(current_roll - 1))
print("losing player score: " + str(losing_player.score))
print("result=" + str(losing_player.score * (current_roll - 1)))
