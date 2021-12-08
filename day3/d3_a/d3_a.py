#!/usr/bin/env python3

input_list = []

#Open input file
with open("input.txt") as f:
    input = f.read()
    input_list = str.splitlines(input)



#total number of binary inputs
total_bins = 0
bin_size = len(input_list[0])
sum_list = [0] * bin_size
gamma_bin_list = [None] * bin_size
eps_bin_list = [None] * bin_size

#Iterate through the input list
for bin_num in input_list:
    for i in range(0, bin_size):
        sum_list[i] += int(list(bin_num)[i])
    total_bins += 1

for i in range(0, bin_size):
    if sum_list[i] <= total_bins/2:
        gamma_bin_list[i] = "0"
        eps_bin_list[i] = "1"
    else: 
        gamma_bin_list[i] = "1"
        eps_bin_list[i] = "0"

gamma_bin = ''.join(gamma_bin_list)
eps_bin = ''.join(eps_bin_list)

gamma_dec = int(gamma_bin, 2)
eps_dec = int(eps_bin, 2)

print("Gamma rate, Binary:{}, Decimal:{}".format(gamma_bin, gamma_dec))
print("Epsilon rate, Binary:{}, Decimal:{}".format(eps_bin, eps_dec))


print("Power Consumption:{}".format(gamma_dec *eps_dec))
