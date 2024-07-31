run:
	go run cmd/server/*.go

build:
	go build -o server cmd/server/main.go

generate:
	templ generate internal/templates/*.templ

clean:
	rm -f server

.PHONY: run build clean
