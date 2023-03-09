# About Allgor
Allgor is an offline data collection and analysis tool.

The MVP is currently running at https://www.allgor.com.br/, written using Vue.js and Firebase.

This repository is the backend for the post-MVP version. See [architecture-overview.png](https://github.com/allgor-data/internal-docs/blob/master/assets/architecture-overview.png)

# Contents
- [How to run](#getting-started-guide)
  - [Development mode](#development-mode)
  - [Docker (Production)](#docker-production)
  - [Building from source](#building-from-source)

## How to run
### Development mode
#### Requirements
- [Go](https://go.dev/) (1.19 or later)
- [Make](https://www.gnu.org/software/make/)
- [Docker Engine](https://docs.docker.com/engine/install/) (or [Docker Desktop](https://docs.docker.com/desktop/install/linux-install/))
- [Docker Compose](https://docs.docker.com/compose/install/)

1. Clone the repository
```sh
$ git clone https://github.com/allgor-data/backend.git allgor-backend
$ cd allgor-backend
```

2. Run the application in development mode (live reload)
```sh
$ docker compose up allgor-backend --playground

# GraphQL server: http://localhost:8080
# GraphQL playground: http://localhost:8080/playground
```

3. (Optional) Run SonarQube locally
```sh
$ docker compose up sonarqube
```

4. (Optional) Generate gqlgen files based on GraphQL schema
```sh
$ make generate
```

5. (Optional) Run all tests
```sh
$ make test
```


### Docker (Production)
Not available yet.

### Building from source
#### Requirements
- [Go](https://go.dev/) (1.19 or later)

1. Using make and Makefile
```sh
$ git clone https://github.com/allgor-data/backend.git allgor-backend
$ cd allgor-backend
$ make build
```

2. Without Makefile
```sh
$ git clone https://github.com/allgor-data/backend.git allgor-backend
$ cd allgor-backend
$ go get -v ./...
$ go build -o ./build/allgor .
```

Now you have the "allgor" binary in `./build/allgor`. If you want to use it globally, run the command below:
```sh
$ sudo mv ./build/allgor /usr/local/bin
```