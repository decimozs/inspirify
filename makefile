go:
	cd utils && go run main.go

build:
	cd utils && go mod tidy

run:
	cd utils && go run main.go
	gcc -o main main.c
	./main
	