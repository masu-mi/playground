CLASSPATH=$(shell echo /usr/local/Cellar/antlr/4.8_1/antlr-4.8-complete.jar:$CLASSPATH)

javac=javac -J-Dfile.encoding=utf-8

.SUFFIXES: .class .java

%.class: %.java
	CLASSPATH=$(CLASSPATH) $(javac) $<

.PHONY: all
all: *.class

%.java: Hello.g4
	antlr Hello.g4

.PHONY: clean

clean:
	rm ./*.java
	rm ./*.class
