build:
	go build -o ./build/generate_pages ./cmd/generate_pages

run:
	./build/generate_pages -output ./gen/
