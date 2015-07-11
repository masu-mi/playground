
all: for-osx.clang.s for-osx.gcc.s for-osx

for-osx.gcc.s: for-osx.c
	gcc-5 ./for-osx.c -S -o for-osx.gcc.s
for-osx.clang.s: for-osx.c
	clang ./for-osx.c -S -o for-osx.clang.s
