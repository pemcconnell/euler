def gcf(ar):
	a = []
	gcf = 1
	for i in ar:
		d = 2
		while not d > i:
			if float(i)/d%1 == 0.0:
				if i == ar[0]:
					a.append(d)
				elif d in a and d > gcf:
					gcf = d
			d += 1
	return gcf

def lcm(ar):
	lcm = ar[0]
	for v in ar[1:]:
		gcd = gcf([lcm,v])
		lcm = (lcm*v)/gcd
	return lcm

print lcm(range(1,21))

