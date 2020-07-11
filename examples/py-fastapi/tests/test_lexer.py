# coding: utf-8

"""
test_lexer
"""

import cryptarithm.lexer
from cryptarithm.lexer import Token

def test_token_eq_as_value():
    assert Token('A') == Token('A')

def test_token_ne_as_value():
    assert not (Token('A') == Token('B'))
    assert Token('A') != Token('B')

def test_tokenize_starndard_expr():
    tokens = cryptarithm.lexer.tokenize('A + B = 10')
    expected = [ Token(s) for s in ['A', '+', 'B', '=', '10', '']]
    assert [tk for tk in tokens] == expected

def test_tokenize_empty_string():
    tokens = cryptarithm.lexer.tokenize('')
    assert [tk for tk in tokens] == [Token('')]

def test_tokenize_whitespaces():
    tokens = cryptarithm.lexer.tokenize('     ')
    assert [tk for tk in tokens] == [Token('')]

def test_tokenize_clogged_expr():
    tokens = cryptarithm.lexer.tokenize('A+B=C')
    expected = [ Token(s) for s in ['A', '+', 'B', '=', 'C', '']]
    assert [tk for tk in tokens] == expected
