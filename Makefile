default: build

install_deps:
	go get -v ./...
	[ -f .env ] || cp .env.example .env

build: install_deps
	go build -o ./build/allgor .

dev:
	go run github.com/cosmtrek/air -c ./.air.toml

clean:
	rm -rf ./build
	rm -rf ./tmp

generate:
	go run github.com/99designs/gqlgen generate

test:
	go test ./...

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out