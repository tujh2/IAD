#!/usr/bin/env python

from contextlib import contextmanager
import time

class cm_timer_2:

	def __enter__(self):
		self.enter_time=time.perf_counter()
		return 333

	def __exit__(self, exp_type, exp_value, traceback):
		print(time.perf_counter()-self.enter_time)




@contextmanager
def cm_timer_1():
	t=time.perf_counter()
	yield 333
	print(time.perf_counter()-t)

if __name__=="main":
	with cm_timer_1():
		time.sleep(2)

	with cm_timer_2():
		time.sleep(2)
