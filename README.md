[![Project Status: Concept – Minimal or no implementation has been done yet, or the repository is only intended to be a limited example, demo, or proof-of-concept.](https://www.repostatus.org/badges/latest/concept.svg)](https://www.repostatus.org/#concept)

# NGINX Unit HTTP Server Template Repository

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](

## What is included on this template?

This template includes all the scaffolding you need to get started on a Go HTTP Server running in NGINX Unit.

## How do I use this template?

**DO NOT FORK** -- this template is meant to be used from the **[`Use this template`](https://github.com/nginxinc/template-repository/generate)** feature.

1. Click on **[`Use this template`](https://github.com/nginxinc/template-repository/generate)**
2. Give a name to your project
3. Wait until the first run of CI finishes (GitHub Actions will process the template and commit to your new repo)
4. Clone your new project and tweak any of the placeholders if necessary. Pay special attention to the README!
5. Happy coding!

**NOTE**: **WAIT** until the first CI run on GitHub Actions finishes before cloning your new project.

---

<!--  DELETE THE LINES ABOVE THIS AND WRITE YOUR PROJECT README BELOW -- PLACEHOLDER SECTIONS HAVE BEEN INCLUDED FOR YOUR CONVENIENCE -->

[![Project Status: Concept – Minimal or no implementation has been done yet, or the repository is only intended to be a limited example, demo, or proof-of-concept.](https://www.repostatus.org/badges/latest/concept.svg)](https://www.repostatus.org/#concept)


# {{REPOSITORY_NAME}}

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

**NOTE**: The WASM support in Go is still experimental and may change in future versions of Go. Additionally, running
WASM in NGINX Unit is currently only supported for C and Rust. As support for Go WASM in NGINX Unit improves, this project 
will be updated to reflect those changes.

### NGINX Unit Configuration

### Prometheus Metrics

The project includes a [Prometheus](https://prometheus.io/) metrics endpoint that can be used to monitor the health of the service.

This can be used as an example to add additional metrics to your project.

The [Prometheus Go client library](https://github.com/prometheus/client_golang) is used to implement the metrics endpoint.

When running, the metrics can be accessed at the following URL:

```shell    
http://<hostname>:<port>/metrics
```

## Getting Started

- Run the tests
- Build the project
- Compile to WebAssembly
- Run the project in NGINX Unit

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

    Please see the [contributing guide](https://github.com/{{REPOSITORY_OWNER}}/{{REPOSITORY_URL}}/blob/main/CONTRIBUTING.md) for guidelines on how to best contribute to this project.

## License

[Apache License, Version 2.0](https://github.com/{{REPOSITORY_OWNER}}/{{REPOSITORY_URL}}/blob/main/LICENSE)

&copy; {{REPOSITORY_OWNER}} 2023