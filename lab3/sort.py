#!/usr/bin/env python

data = [4, -30, 100, -100, 123, 1, 0, -1, -4]

if __name__ == '__main__':
	result = sorted(data)
	print(result)

	result_with_lambda = lambda data: sorted(data)
	print(result_with_lambda(data))
