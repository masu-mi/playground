# coding: utf-8

import copy
import random
from typing import Any,Callable,List
from .card import Card,Suit,simple_index

class Deck():
    cards: List[Card]

    def __init__(self, cards: List[Card] = []):
        self.cards = copy.deepcopy(cards)

    def push(self, c: Card) -> None:
        self.cards.append(c)
    def pop(self) -> Card:
        return self.cards.pop()
    def shuffle(self) -> None:
        random.shuffle(self.cards)
    def sort(self, comp: Callable[[Card], Any] = simple_index) -> None:
        random.shuffle(self.cards)
