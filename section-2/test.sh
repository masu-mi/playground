#!/usr/bin/env bash
(cd ./parsing/lexer/;javac Test.java;java Test [a,b,c,[a,	b,  hoge,[fuga,   fuga]]]; java Test [a,b [c]fda)
(cd ./parsing/recursive-descent;javac -cp . Test.java;java -cp . Test '[a,b,c,[a,	b,  hoge,[fuga,   fuga]]]'; java -cp . Test '[a,b [c]')
