with open('AoCDay9-data', 'r') as fh:
    gameInput = fh.read()

gameInput = '18,1,0,161,255,137,254,252,14,95,165,33,181,168,2,188'
# gameInput = [int(x) for x in gameInput.split(',')]
gameInput = [ord(c) for c in gameInput] + [17, 31, 73, 47, 23]
skip = 0
idx = 0
l = list(range(256))

for i in range(64):
    for length in gameInput:
        if idx + length >= len(l):
            pivot = (idx+length)% len(l)
            temp = l[idx:] + l[:pivot]
            temp = list(reversed(temp))
            l[idx:], l[:pivot] = temp[:length-pivot],temp[length-pivot:]
        else:
            temp = l[idx:idx+length]
            temp = list(reversed(temp))
            l[idx:idx + length] = temp
        idx = (idx+length+skip)%len(l)
        skip += 1

        print(idx, skip)

hash = [0]*16
for i in range(16):
    val =0
    for j in range(16):
        val ^= l[i*16+j]
    hash[i] = val

print(hash)
print(''.join("{:02x}".format(x) for x in hash))
print(l[0]*l[1])
