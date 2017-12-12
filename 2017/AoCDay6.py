

import numpy as np

gameInput = "2	8	8	5	4	2	3	1	5	5	1	2	15	13	5	14"
gameInput = tuple([int(i) for i in gameInput.split()])

def redistribute(buckets):
    #findmax
    buckets = list(buckets)
    m = max(buckets)
    for i,v in enumerate(buckets):
        if v == m:
            break
    buckets[i] = 0

    for j in range(v):
        index = (i+j+1)%len(buckets)
        buckets[index] += 1

    # print(buckets)
    return tuple(buckets)

steps = 0
history = set()
buckets = gameInput
while buckets not in history:
    history.add(buckets)
    steps += 1
    buckets = redistribute(buckets)
print("finished {}".format(steps))

end = buckets
steps2 = 1
buckets = redistribute(buckets)
while buckets != end:
    steps2+=1
    buckets = redistribute(buckets)

print(steps2)