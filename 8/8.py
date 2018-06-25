from operator import mul
import re

size = 5
s = open('8.txt', 'r').read().replace('\n', '')

ar = []
i = size
while i > 0:
	ar += re.findall(r'\d{' + str(size) + '}', s[i:])
	i -= 1

m = 0
for n in ar:
	n = reduce(mul,map(int, str(n)),1)
	if n > m:
		m = n

print m
