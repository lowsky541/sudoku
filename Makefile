
build: $(wildcard *.go)
	mkdir -p bin
	go build -o bin/sudoku $?
