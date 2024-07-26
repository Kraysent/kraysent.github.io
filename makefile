.PHONY: build run

build:
	go build -o ./build/generate_pages ./cmd/generate_pages

run: build
	./build/generate_pages -output ./gen/

clean:
	rm -rf ./gen/
	rm -rf ./build/
