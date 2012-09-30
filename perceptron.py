#!/usr/bin/env python
# -*- coding: utf-8 -*-
import sys
from numpy import *

def predict(weight, input):
  return copysign(1.0, dot(weight, input))

def train(weight, input, ans):
  output = predict(weight, input)
  if (output * ans < 0):
    return weight + ans * input
  else:
    return weight


input_list = [
    (array([255,   0,   0]),  1.0),
    (array([  0, 255, 255]), -1.0),
    (array([  0, 255,   0]), -1.0),
    (array([255,   0, 255]),  1.0),
    (array([  0,   0, 255]), -1.0),
    (array([255, 255,   0]),  1.0)
    ]

# 訓練パート
weight = array([1, 0, 0])
for input_data in input_list:
  weight = train(weight, input_data[0], input_data[1])

# 推定パート
for line in sys.stdin:
  input = line.split()
  output = predict(weight, array([
              float(input[0]),
              float(input[1]),
              float(input[2])
            ]))
  if (output > 0):
    print "warm color"
  else:
    print "cold color"
