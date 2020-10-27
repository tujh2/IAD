#!/usr/bin/env python

class Unique(object):
	def __init__(self, items, **kwargs):
		self.items = list(items)
		self.uniqItems = list()
		self.index = 0
		try:
			self.ignore_case = bool(kwargs['ignore_case'])
		except (KeyError, TypeError):
			self.ignore_case = False

	def __next__(self):
		while True:
			if self.index >= len(self.items):
				raise StopIteration
			else:
				current = self.items[self.index]
				if (self.ignore_case == True) & (type(current) == str):
					current = current.upper()
				if current not in self.uniqItems:
					self.uniqItems.append(current)
					return self.items[self.index]
				self.index = self.index + 1

	def __iter__(self):
		return self
