[![Project Status: Concept – Minimal or no implementation has been done yet, or the repository is only intended to be a limited example, demo, or proof-of-concept.](https://www.repostatus.org/badges/latest/concept.svg)](https://www.repostatus.org/#concept)

# unit-http-server-template

This project is intended to serve three purposes:
- Provide a template for creating HTTP Servers in Go
- Provide a reference for how to create a Go project that can be compiled to [WebAssembly](https://webassembly.org/);
- Provide a reference for how to create a Go project that can be compiled to WebAssembly and run in [NGINX Unit](https://unit.nginx.org/); 

## Requirements

This project requires the following tools be installed on your system:
- [Go](https://golang.org/), version 1.21 or higher

Included is a [asdf](https://asdf-vm.com/#/) configuration file that will install the required version of Go if you have asdf installed, using the following command:

```shell
asdf install
```

## Features

- WebAssembly compilation
- NGINX Unit configuration
- Prometheus metrics

### WebAssembly Compilation

Compilation to WebAssembly uses the [WASI support](https://go.dev/blog/wasi) in Go 1.21. 

### NGINX Unit Configuration

### Prometheus Metrics

The project includes a [Prometheus](https://prometheus.io/) metrics endpoint that can be used to monitor the health of the service.

This can be used as an example to add additional metrics to your project.

The [Prometheus Go client library](github.com/prometheus/client_golang) is used to implement the metrics endpoint.

When running, the metrics can be accessed at the following URL:

```shell    
http://<hostname>:<port>/metrics
```

## Getting Started

- Use the template repository to create a new repository
- Run the tests
- Build the project
- Compile to WebAssembly
- Run the project in NGINX Unit

### Use the Template Repository

To use this template repository, [generate a new repository with the template](https://github.com/ciroque/go-base-http-server/generate).

### Run the Tests

To run the tests, execute the following command:

```shell
go test ./...
``` 

### Build the Project

To build the project, execute the following command:

```shell    
go build -o _build/http-server cmd/main/main.go
```

### Compile to WebAssembly

To compile the project to WebAssembly, execute the following command:

```shell
GOOS=js GOARCH=wasm go build -o _build/http-server.wasm cmd/main/main.go
```

```shell
GOOS=wasip1 GOARCH=wasm go build -o _build/http-server.wasm cmd/main/main.go
```

```shell
GOOS=wasip2 GOARCH=wasm go build -o _build/http-server.wasm cmd/main/main.go
```

### Run the Project in NGINX Unit

The easiest way to run the project in NGINX Unit is to use Docker. The `Dockerfile` can be used to build an
image that can be used to run the project in NGINX Unit.

To build the image, execute the following command:

```shell    
docker build -t bhs:dev .
```

To run the image, execute the following command:

```shell
docker run -dit -p 8888:8888 --rm --name bhs_unit_dev bhs:dev
```

Throw some requests at the server:

```shell
curl http://localhost:8888
curl http://localhost:8888/metrics
```

To stop the container, execute the following command:

```shell
docker stop bhs_unit_dev
```

To remove the image, execute the following command:

```shell
docker rmi bhs:dev
```




## Contributing

Please see the [contributing guide](https://github.com/ciroque/go-base-http-server/blob/main/CONTRIBUTING.md) for guidelines on how to best contribute to this project.

## License

[Apache License, Version 2.0](https://github.com/ciroque/go-base-http-server/blob/main/LICENSE)

&copy; Steve Wagner [ciroque](https://github.com/ciroque) 2023