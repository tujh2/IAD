#!/usr/bin/env python

import json
import cm_timer
import field
import gen_random
import print_result
import unique

@print_result.print_result
def f1(arg):
	return sorted(unique.Unique(list(field.field(arg, "job-name")), ignore_case = True), key=str.lower)


@print_result.print_result
def f2(arg):
	return list(filter(lambda x: x.lower().startswith("программист"), arg))


@print_result.print_result
def f3(arg):
	return list(map(lambda x: x + " с опытом Python", arg))


@print_result.print_result
def f4(arg):
	return list(zip(arg, map(lambda x: "зарплата " + str(x), gen_random.gen_random(len(arg), 100000, 200000))))


path = "data_light.json"

if __name__ == '__main__':
	with cm_timer.cm_timer_1():
		with open(path) as f:
			data = json.load(f)
			f4(f3(f2(f1(data))))

