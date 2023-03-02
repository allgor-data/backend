default: build

install_deps:
	go get -v ./...

build: install_deps
	go build -o ./build/allgor .

dev:
	go run github.com/cosmtrek/air -c ./.air.toml

clean:
	rm -rf ./build

generate:
	go run github.com/99designs/gqlgen generate