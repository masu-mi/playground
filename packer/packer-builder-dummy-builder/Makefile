
.PHONNY: install build all

build: packer-builder-dummy-builder

packer-builder-dummy-builder: main.go
	go build .

install: packer-builder-dummy-builder
	cp $< $(HOME)/.packer.d/plugins/$<
