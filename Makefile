VERSION = `date +%y.%m`

ifndef VERSION
    VERSION = "n/a"
endif


.PHONY: all clean build install uninstall


all: clean build

build:
	@echo 'Building roguelike...'
	@go build -ldflags '-s -w -X main.Version='${VERSION}

clean:
	@echo 'Cleaning...'
	@go clean

install: build
	@echo installing executable file to /usr/bin/roguelike
	@sudo cp go_roguelike /usr/bin/roguelike

uninstall: clean
	@echo removing executable file from /usr/bin/roguelike
	@sudo rm /usr/bin/roguelike
