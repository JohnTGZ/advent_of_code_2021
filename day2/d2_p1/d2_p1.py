#!/usr/bin/env python3

cmd_list = []

#Open input file
with open("input.txt") as f:
    input = f.read()
    cmd_list = str.splitlines(input)

def cmd_to_idx(command):
    return {    
        "forward": 0,
        "down": 1,
        "up": 2,
    }[command]


movement_culmulative = [0, 0]
aim = 0

#Iterate through the commands
for cmd_full in cmd_list:
    #separate cmd from value
    cmd_full_split = cmd_full.split(' ', 1)
    cmd = cmd_full_split[0]
    value = int(cmd_full_split[1])

    if cmd_to_idx(cmd) == 0:
        #forward

        #Increase horizontal position by x 
        movement_culmulative[0] += value

        #Increase depth by aim multiplied by x
        movement_culmulative[1] += aim*value
        
    elif cmd_to_idx(cmd) == 1:
        #down

        #Increase depth
        # movement_culmulative[1] += value

        #Increases aim by x
        aim += value
    else:
        #up

        #Decrease depth
        # movement_culmulative[1] -= value

        #Decreases aim by x
        aim -= value



print("Total forward: {}".format(movement_culmulative[0]))
print("Total depth: {}".format(movement_culmulative[1]))


print("Answer: {}".format(movement_culmulative[1]*movement_culmulative[0]))
