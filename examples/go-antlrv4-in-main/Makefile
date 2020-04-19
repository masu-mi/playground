.PHONY: all build

all: build

build: example

example: calc_parser.go main.go
	go build

calc_%.go: Calc.g4
	# antlr  -Dlanguage=Go -package main $<
	antlr -listener -visitor -Dlanguage=Go -package main $<
