

with open('AoCDay8-data') as fh:
    gameInput = [line.split() for line in fh]
    gameInput = [(x[0], "{} {} {}".format(x[0], "+=" if x[1] =='inc' else "-=", x[2]), " ".join(x[4:])) for x in gameInput]

# set all registers to 0
for instruction in gameInput:
    exec("{} = 0".format(instruction[0]))

max_value = 0

for instruction in gameInput:
    exec("cond = {}".format(instruction[2]))
    if cond:
        exec(instruction[1])
    exec("cond = {}>max_value".format(instruction[0]))
    if cond:
        exec("max_value = {}".format(instruction[0]))

results =[]
for instruction in gameInput:
    exec("results.append({})".format(instruction[0]))
    # results.append(exec(instruction[0]))

print(max(results))
print(max_value)