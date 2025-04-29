##
# eol
#
# @file
# @version 0.1

name = eol
PREFIX = /usr/local

all: build

build:
	go build -o $(name)

install: all
	cp -f $(name) "$(DESTDIR)$(PREFIX)/bin/"
	chmod 755 "$(DESTDIR)$(PREFIX)/bin/$(name)"

uninstall:
	rm -f "$(DESTDIR)$(PREFIX)/bin/$(name)"

clean:
	rm -f $(name)

# end
