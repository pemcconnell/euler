import string

s = open('13.txt', 'r').read().strip()

print str(sum(int(c) for c in string.split(s, "\n")))[:10]
