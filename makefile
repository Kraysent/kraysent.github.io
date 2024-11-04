.PHONY: build run

include .env

build:
	go build -o ./build/generate_pages ./cmd/generate_pages

run: build
	./build/generate_pages -output ./gen/
	cp -r static/* gen/

clean:
	rm -rf ./gen/
	rm -rf ./build/

install-jekyll:
	cd gen && bundle install

serve:
	cd gen && bundle exec jekyll serve
