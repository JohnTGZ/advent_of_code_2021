#!/usr/bin/env python3

number_list = []

with open("../input.txt") as f:
    input = f.read()
    number_list = str.splitlines(input)

#If next number is higher than current number 
#then increment deeper_counter 
deeper_counter = 0
not_deeper_counter = 0


#iterate through the list
#start from second number 
for i in range(0, len(number_list)-1):
    if (int(number_list[i+1]) >=  int(number_list[i])):
        print("{} >= {}".format(number_list[i+1], number_list[i]))
        deeper_counter = deeper_counter + 1
    else:
        # print("{} < {}".format(number_list[i+1], number_list[i]))
        not_deeper_counter += 1


print("Measurements larger than prev measurements: {}".format(deeper_counter))
print("Not deeper: {}".format(not_deeper_counter))
