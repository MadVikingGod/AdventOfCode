

class State(object):
    groups = 0
    levels = 1
    garbage = 0
    is_garbage = False
    is_ignore = False

    def read_char(self, char):
        if self.is_ignore:
            self.is_ignore = False
            return
        if self.is_garbage:
            if char == '>':
                self.is_garbage = False
                return
            if char == '!':
                self.is_ignore = True
                return
            self.garbage += 1
        else:
            if char == '{':
                self.groups += self.levels
                self.levels += 1
            if char == '<':
                self.is_garbage = True
            if char == '}':
                self.levels -= 1

if __name__ == "__main__":
    with open('AoCDay9-data', 'r') as fh:
        gameInput = fh.read()

    state = State()
    for char in gameInput:
        state.read_char(char)

    print(state.groups)
    print(state.garbage)