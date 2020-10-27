#!/usr/bin/env python

def print_result(func_to_decorate):
	def decorated_func(*args,**kwargs):
		a = func_to_decorate(*args,**kwargs)
		print(func_to_decorate.__name__)
		if type(a) == list:
			for i in a:
				print(i)
			print("\n")
		elif type(a) == dict:
			for i, j in dict(a.items()):
				print(i, "=", j)
			print("\n")
		else:
			print(a,"\n")
		return a
		
	return decorated_func
