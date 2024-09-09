build:
	go build -o bin/redirect

run: build
	./bin/redirect
