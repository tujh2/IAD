#!/usr/bin/env python
from __future__ import annotations
from abc import ABC, abstractmethod
from random import randrange
from typing import List


class Subject(ABC):

    @abstractmethod
    def attach(self, observer: Observer) -> None:
        pass

    @abstractmethod
    def detach(self, observer: Observer) -> None:
        pass

    @abstractmethod
    def notify(self) -> None:
        pass


class SomeContentManager(Subject):

    _state: int = None
    _observers: List[Observer] = []

    def attach(self, observer: Observer) -> None:
        print("SomeContentManager: Attached an observer.")
        self._observers.append(observer)

    def detach(self, observer: Observer) -> None:
        self._observers.remove(observer)

    def notify(self) -> None:
        print("SomeContentManager: Notifying observers...")
        for observer in self._observers:
            observer.update(self)

    def request_content(self) -> None:

        print("\nGoing to the internet(not) and getting the data.")
        self._state = randrange(0, 10)

        self.notify()


class Observer(ABC):

    @abstractmethod
    def update(self, subject: Subject) -> None:
        pass



class MyObserverA(Observer):
    def update(self, subject: Subject) -> None:
        print("MyObserverA: update {}".format(subject._state))


class MyObserverB(Observer):
    def update(self, subject: Subject) -> None:
        print("MyObserverA: update {}".format(subject._state))


if __name__ == "__main__":
    content_manager = SomeContentManager()

    observer_a = MyObserverA()
    content_manager.attach(observer_a)

    observer_b = MyObserverB()
    content_manager.attach(observer_b)

    content_manager.request_content()
    content_manager.request_content()

    content_manager.detach(observer_a)
    content_manager.request_content()
    content_manager.detach(observer_b)
    content_manager.request_content()


