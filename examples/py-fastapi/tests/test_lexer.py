# coding: utf-8

"""
test_lexer
"""

from cryptarithm.lexer import Token

def test_token_eq_as_value():
    assert Token('A') == Token('A')

def test_token_ne_as_value():
    assert not (Token('A') == Token('B'))
    assert Token('A') != Token('B')

def test_tokenize_starndard_expr():
    tokens = Token.tokenize('A + B = 10')
    expected = [ Token(s) for s in ['A', '+', 'B', '=', '10']]
    assert [tk for tk in tokens] == expected

def test_tokenize_empty_string():
    tokens = Token.tokenize('')
    assert [tk for tk in tokens] == []

def test_tokenize_whitespaces():
    tokens = Token.tokenize('     ')
    assert [tk for tk in tokens] == []

def test_tokenize_clogged_expr():
    tokens = Token.tokenize('A+B=C')
    expected = [ Token(s) for s in ['A', '+', 'B', '=', 'C']]
    assert [tk for tk in tokens] == expected
