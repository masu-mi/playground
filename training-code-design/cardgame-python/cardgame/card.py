# coding: utf-8
# vim: fileencoding=utf-8

from enum import Enum
from typing import Any,NamedTuple,Optional

class Suit(Enum):
    CLUBS = 1
    DIAMONDS = 2
    HEARTS = 3
    SPADES = 4

class Card(NamedTuple):
    suit: int
    number: int
    special_name: Optional[str] = None

def simple_index(c: Card) -> Any:
    if c.suit == 0:
        return c.special_name
    return c.number*4+c.suit
