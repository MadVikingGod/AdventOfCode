class CircularList(object):
    def __init__(self, l=[]):
        self.l = list(l)

    def __getitem__(self, item):
        length = len(self.l)
        if isinstance(item, int):
            return self.l[item % length]
        elif isinstance(item, slice):
            slice_length = item.stop - item.start
            start = item.start % length
            if slice_length + start < length:
                return self.l[item.start % length: item.stop % length]
            upper_slice = self.l[item.start%length:]
            return upper_slice + self.l[:slice_length-len(upper_slice)]

    def __setitem__(self, key, value):
        length = len(self.l)
        if isinstance(key, int):
            self.l[key % length] = value
        elif isinstance(key, slice):
            slice_length = key.stop - key.start
            start = key.start % length
            if slice_length + start < length:
                self.l[key.start % length: key.stop % length] = value
                return
            upper_len = length-key.start
            self.l[key.start % length:] = value[:upper_len]
            remaining = slice_length - upper_len
            self.l[:remaining] = value[upper_len:]

    def __repr__(self):
        return repr(self.l)
    def __str__(self):
        return str(self.l)


if __name__ == "__main__":
    gameInput = '18,1,0,161,255,137,254,252,14,95,165,33,181,168,2,188'
    gameInput = [int(x) for x in gameInput.split(',')]

    l = CircularList(range(256))
    skip = 0
    idx = 0

    for length in gameInput:
        if length >0:
            l[idx:idx+length] = list(reversed(l[idx:idx+length]))
        idx = (idx + skip + length) % 256
        skip +=1

    print(l[0]*l[1])

    gameInput = '18,1,0,161,255,137,254,252,14,95,165,33,181,168,2,188'
    gameInput = [ord(c) for c in gameInput] + [17, 31, 73, 47, 23]

    l = CircularList(range(256))
    skip = 0
    idx = 0

    for i in range(64):
        for length in gameInput:
            if length >0:
                l[idx:idx+length] = list(reversed(l[idx:idx+length]))
            idx = (idx + skip + length) % 256
            skip +=1

    hash = [0] * 16
    for i in range(16):
        val = 0
        for j in range(16):
            val ^= l[i * 16 + j]
        hash[i] = val

    print(hash)
    print(''.join("{:02x}".format(x) for x in hash))