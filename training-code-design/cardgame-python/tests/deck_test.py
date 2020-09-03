# coding: utf-8
import random
from cardgame import card
from cardgame import deck

def test_deck_shuffle():
    random.seed(0)
    d = deck.Deck()

    inputs = [ card.Card(suit = card.Suit.SPADES, number = num) for num in range(1, 13) ]
    for c in inputs:
        d.push(c)
    assert d.cards == inputs
    d.shuffle()
    expected = [ card.Card(suit = card.Suit.SPADES, number = num)
                for num in [2, 10, 9, 6, 11, 3, 4, 8, 5, 1, 12, 7] ]
    assert d.cards == expected
