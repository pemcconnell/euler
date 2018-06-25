def sumn(n):
	i = 1
	r = 0
	rs = 0
	while i <= n:
		r += i
		rs += i*i
		i += 1
	return [r,rs]

sn = sumn(100)
print (sn[0]*sn[0])-sn[1]
