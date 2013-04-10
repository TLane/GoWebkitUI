all: build

build: mkbin
	go build -o bin/GoWebkitUI WebkitUI.go

buildOptimized: mkbin
	go build -o bin/GoWebkitUI WebkitUI.go -gcflags -O3

mkbin:
	test -d bin || mkdir bin

clean:
	rm bin/*
