GOPATH	= $(CURDIR)
BINDIR	= $(CURDIR)/bin

PROGRAMS = mqtt-forwarder

.PHONY: build
build:
	env GOPATH=$(GOPATH) go install $(PROGRAMS)

.PHONY: destdirs
destdirs:
	mkdir -p -m 0755 $(DESTDIR)/usr/bin

.PHONY: strip
strip: build
	strip --strip-all $(BINDIR)/mqtt-forwarder

.PHONY: install
install: strip destdirs install-bin

.PHONY: install-bin
install-bin:
	install -m 0755 $(BINDIR)/mqtt-forwarder $(DESTDIR)/usr/bin

.PHONY: clean
clean:
	/bin/rm -f bin/mqtt-forwarder

.PHONY: uninstall
uninstall:
	/bin/rm -f $(DESTDIR)/usr/bin

.PHONY: depend
depend:
	env GOPATH=$(GOPATH) go get -u github.com/sirupsen/logrus/
	env GOPATH=$(GOPATH) go get -u github.com/eclipse/paho.mqtt.golang
	env GOPATH=$(GOPATH) go get -u gopkg.in/ini.v1
	env GOPATH=$(GOPATH) go get -u github.com/nu7hatch/gouuid

.PHONY: distclean
distclean:
	/bin/rm -rf src/github.com/
	/bin/rm -rf src/gopkg.in/

.PHONY: all
all: build strip install

