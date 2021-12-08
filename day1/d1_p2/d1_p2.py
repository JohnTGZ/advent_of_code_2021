#!/usr/bin/env python3

number_list = []

with open("../input.txt") as f:
    input = f.read()
    number_list = str.splitlines(input)

#If next number is higher than current number 
#then increment deeper_counter 
deeper_counter = 0
not_deeper_counter = 0


def sum(list):
    sum_result = 0
    for elem in list:
        sum_result += int(elem)
    return int(sum_result)

#iterate through the list
#start from second number 
for i in range(2, len(number_list)-1):
    window_1 = sum(number_list[i-2:i+1])
    window_2 = sum(number_list[i-1:i+2])
    if (window_2 > window_1):
        print("{} > {}".format(window_2, window_1))
        deeper_counter = deeper_counter + 1
    else:
        # print("{} < {}".format(number_list[i+1], number_list[i]))
        not_deeper_counter += 1


print("Measurements larger than prev measurements: {}".format(deeper_counter))
print("Not deeper: {}".format(not_deeper_counter))
