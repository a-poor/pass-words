.PHONY: build
build:
	@echo "Building..."
	go build -o ./build/pass-words pass-words.go

run: 
	@echo "Running..."
	go run pass-words.go

	