generate:
	templ generate

run: generate
	go run .

.PHONY: generate run
