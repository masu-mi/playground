tmp := tmp
files := files

LIBS = -lssl -lcrypto
TARGET := echo_raw echo_ssl_bio echo_ssl_low

.PHONY: build setup
build: $(addprefix $(tmp)/,$(TARGET))

$(tmp)/echo_raw: LIBS =
$(tmp)/%: %.c $(tmp)
	gcc -Wall -o $@ $< $(LIBS)

setup: $(files)
	openssl genrsa 4096 > $(files)/key.pem
	openssl req -new -key $(files)/key.pem > $(files)/csr.pem
	openssl x509 -req -signkey $(files)/key.pem < $(files)/csr.pem > $(files)/cert.pem

$(files):
	mkdir $@
$(tmp):
	mkdir $@
