import re

SIZE = 4
GREATESTPRODUCT = 0

def find_biggest(cells):
	global GREATESTPRODUCT
	gp = 0
	l = len(cells)
	for i,num in enumerate(cells):
		if i > (l - SIZE):
			break
		x = SIZE-1
		p = 1
		while x >= 0:
			p *= int(cells[i+x])
			x -= 1
		if p > gp:
			gp = p
	if gp > GREATESTPRODUCT:
		GREATESTPRODUCT = gp

data = []
cl = 0
lines = open('11.txt', 'r').read().split('\n')
for line in lines:
	cells = line.split(' ')
	if cl == 0:
		cl = len(cells)
	data.append(cells)
	# left/right
	find_biggest(cells)

diagonals = []
for i,line in enumerate(lines):
	x = 0
	v = []
	dr = []
	dl = []
	while x < cl:
		v.append(data[x][i])
		if not dr == False:
			try:
				dr.append(data[x][i+x])
			except IndexError:
				if len(dr) < SIZE:
					dr = False
		if not dl == False:
			try:
				dl.append(data[x][i-x])
			except IndexError:
				if len(dl) < SIZE:
					dl = False
		x += 1
	if not dr == False:
		diagonals += dr
	if not dl == False:
		diagonals += dl
	#up/down
	find_biggest(v)
	#diagonals
	find_biggest(diagonals)

print GREATESTPRODUCT
