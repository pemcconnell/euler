t = 10001

def findprime(t):
	pc = 0
	t -= 1
	num = 3
	while True:
		for i in range(2,num):
			if num%i == 0:
				break
		else:
			pc += 1
			if pc == t:
				return num
		num += 2

print findprime(t)
