

with open('AoCDay5-data') as fh:
    gameInput = [int(line) for line in fh]


print(len(gameInput))
# gameInput = [0, 3, 0, 1, -3]

ip = 0
steps = 0
while ip < len(gameInput):
    ip, old_ip = gameInput[ip] + ip, ip
    if gameInput[old_ip] >= 3:
        gameInput[old_ip] -= 1
    else:
        gameInput[old_ip] += 1
    steps +=1

print(steps)