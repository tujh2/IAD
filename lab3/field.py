#!/usr/bin/env python

def field(items, *args):
	assert len(args) > 0
	for i in items:
		for j in args:
			yield i[j]
