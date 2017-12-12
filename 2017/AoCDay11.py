from collections import Counter
with open('AoCDay11-data', 'r') as fh:
    gameInput = fh.read()

x = Counter(gameInput)

n=0
s=0
w=0
e=0
m =0
for step in gameInput.split(','):
    if 'n' in step:
        n+=1
    if 's' in step:
        s += 1
    if 'e' in step:
        e += 1
    if 'w' in step:
        w += 1
    if m < max(abs(n-s),abs(e-w)):
        m = max(abs(n-s),abs(e-w))

print(m)
